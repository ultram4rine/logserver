package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ultram4rine/logserver/proto"

	"github.com/BurntSushi/toml"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/alecthomas/kingpin.v2"
)

var config struct {
	Port string `toml:"listen_port"`
	DB   db     `toml:"db"`
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

	if _, err := toml.DecodeFile(*confpath, &config); err != nil {
		log.Fatalf("Error decoding config file from %s", *confpath)
	}

	conn, err := sqlx.Connect("clickhouse", fmt.Sprintf("%s?username=%s&password=%s&database=%s", config.DB.Host, config.DB.User, config.DB.Pass, config.DB.Name))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	defer conn.Close()

	listener, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterLogServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) GetAvailableSwitches(c context.Context, request *pb.SwName) (response *pb.Switch, err error) {
	return nil, nil
}

func (s *server) GetDHCPLog(c context.Context, request *pb.DHCPLogEntry) (response *pb.DHCPLog, err error) {
	return nil, nil
}

func (s *server) GetSwitchLog(c context.Context, request *pb.SwitchLogEntry) (response *pb.SwitchLog, err error) {
	return nil, nil
}
