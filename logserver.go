package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.sgu.ru/ultramarine/logserver/pb"
	"git.sgu.ru/ultramarine/logserver/service"

	"github.com/BurntSushi/toml"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
)

var confPath = kingpin.Flag("conf", "Path to config file.").Short('c').Default("logserver.conf.toml").String()

var conf struct {
	App app `toml:"app"`
	DB  db  `toml:"db"`
}

func main() {
	kingpin.Parse()

	if _, err := toml.DecodeFile(*confPath, &conf); err != nil {
		log.Fatalf("Failed to decode config file from %s", *confPath)
	}

	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "clickhouse", fmt.Sprintf("%s?username=%s&password=%s&database=%s", conf.DB.Host, conf.DB.User, conf.DB.Pass, conf.DB.Name))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Info("Connected to ClickHouse database")

	var (
		svc     = service.LogService{DB: db}
		errChan = make(chan error, 1000)
	)

	go func() {
		listener, err := net.Listen("tcp", ":"+conf.App.ListenPort)
		if err != nil {
			errChan <- err
			return
		}

		creds, err := credentials.NewServerTLSFromFile(conf.App.CertPath, conf.App.KeyPath)
		if err != nil {
			errChan <- err
			return
		}

		gRPCServer := grpc.NewServer(grpc.Creds(creds))
		pb.RegisterLogServiceServer(gRPCServer, svc)

		log.Infof("Started LogServer on %s port", conf.App.ListenPort)

		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		b, err := ioutil.ReadFile(conf.App.ClientCertPath)
		if err != nil {
			errChan <- err
			return
		}

		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			errChan <- errors.New("Failed to append certificates")
			return
		}

		tlsConfig := &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            cp,
		}

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))}
		err = pb.RegisterLogServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", conf.App.ListenPort), opts)
		if err != nil {
			errChan <- err
			return
		}

		http.ListenAndServe(":"+conf.App.GatewayPort, mux)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%v", <-c)
	}()

	log.Info(<-errChan)
}

type app struct {
	CertPath       string `toml:"cert_path"`
	KeyPath        string `toml:"key_path"`
	ClientCertPath string `toml:"client_cert_path"`
	ListenPort     string `toml:"listen_port"`
	GatewayPort    string `toml:"gateway_port"`
}

type db struct {
	Host string `toml:"host"`
	Name string `toml:"name"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}
