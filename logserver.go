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

	"git.sgu.ru/ultramarine/logserver/auth"
	"git.sgu.ru/ultramarine/logserver/conf"
	"git.sgu.ru/ultramarine/logserver/pb"
	"git.sgu.ru/ultramarine/logserver/service"

	_ "github.com/ClickHouse/clickhouse-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
)

var confPath = kingpin.Flag("conf", "Path to config file.").Short('c').Default("logserver.conf.toml").String()

func main() {
	kingpin.Parse()

	if err := conf.ParseConfig(*confPath); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "clickhouse", fmt.Sprintf("%s?username=%s&password=%s&database=%s", conf.Conf.DB.Host, conf.Conf.DB.User, conf.Conf.DB.Pass, conf.Conf.DB.Name))
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
		listener, err := net.Listen("tcp", ":"+conf.Conf.App.ListenPort)
		if err != nil {
			errChan <- err
			return
		}

		creds, err := credentials.NewServerTLSFromFile(conf.Conf.App.CertPath, conf.Conf.App.KeyPath)
		if err != nil {
			errChan <- err
			return
		}

		gRPCServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.LDAPAuthFunc),
		)))
		pb.RegisterLogServiceServer(gRPCServer, svc)

		log.Infof("Started LogServer on %s port", conf.Conf.App.ListenPort)

		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		b, err := ioutil.ReadFile(conf.Conf.App.ClientCertPath)
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
		err = pb.RegisterLogServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", conf.Conf.App.ListenPort), opts)
		if err != nil {
			errChan <- err
			return
		}

		http.ListenAndServe(":"+conf.Conf.App.GatewayPort, mux)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%v", <-c)
	}()

	log.Info(<-errChan)
}
