package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.sgu.ru/ultramarine/logserver"
	pb "git.sgu.ru/ultramarine/logserver/logpb"

	"github.com/BurntSushi/toml"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
)

var conf struct {
	App app `toml:"app"`
	DB  db  `toml:"db"`
}

type app struct {
	CertPath string `toml:"cert_path"`
	KeyPath  string `toml:"key_path"`
	Port     string `toml:"listen_port"`
}

type db struct {
	Host string `toml:"host"`
	Name string `toml:"name"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

var confpath = kingpin.Flag("conf", "Path to config file.").Short('c').Default("logserver.conf.toml").String()

func main() {
	kingpin.Parse()

	if _, err := toml.DecodeFile(*confpath, &conf); err != nil {
		log.Fatalf("Error decoding config file from %s", *confpath)
	}

	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "clickhouse", fmt.Sprintf("%s?username=%s&password=%s&database=%s", conf.DB.Host, conf.DB.User, conf.DB.Pass, conf.DB.Name))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	defer db.Close()

	var svc logserver.Service
	svc = logserver.LogService{DB: db}
	errChan := make(chan error, 1000)

	endpoints := logserver.Endpoints{
		DHCPEndpoint:    logserver.MakeDHCPEndpoint(svc),
		SwitchEndpoint:  logserver.MakeSwitchEndpoint(svc),
		SimilarEndpoint: logserver.MakeSimilarEndpoint(svc),
	}

	go func() {
		listener, err := net.Listen("tcp", ":"+conf.App.Port)
		if err != nil {
			errChan <- err
			return
		}

		creds, err := credentials.NewServerTLSFromFile(conf.App.CertPath, conf.App.KeyPath)
		if err != nil {
			errChan <- err
			return
		}

		handler := logserver.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer(grpc.Creds(creds))
		pb.RegisterLogServer(gRPCServer, handler)

		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Println(<-errChan)
}
