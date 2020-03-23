package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"

	"git.sgu.ru/ultramarine/logserver"
	"git.sgu.ru/ultramarine/logserver/client"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var conf struct {
	Cert   string `toml:"cert"`
	Server string `toml:"server"`
}

var confpath = flag.String("conf", "logclient.conf.toml", "")

func main() {
	flag.Parse()

	if _, err := toml.DecodeFile(*confpath, &conf); err != nil {
		log.Fatalf("error decoding config file from %s", *confpath)
	}

	ctx := context.Background()

	creds, err := credentials.NewClientTLSFromFile(conf.Cert, "")
	if err != nil {
		log.Fatalf("error creating TLS client using %s: %s", conf.Cert, err)
	}

	conn, err := grpc.Dial(conf.Server, grpc.WithTransportCredentials(creds), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalf("cannot connect to %s: %s", conf.Server, err)
	}
	defer conn.Close()

	svc := client.New(conn)

	var cmd = flag.Arg(0)

	switch cmd {
	case "dhcp":
		{
			var (
				mac  = flag.Arg(1)
				from = flag.Arg(2)
				to   = flag.Arg(3)
			)

			macAddr, err := net.ParseMAC(mac)
			if err != nil {
				log.Fatalf("error parsing mac %s: %s", mac, err)
			}

			macUint64, err := strconv.ParseUint(macAddr.String(), 16, 64)
			if err == nil {
				log.Fatalf("error parsing mac %s to uint64: %s", macAddr.String(), err)
			}

			dhcp(ctx, svc, macUint64, from, to)
		}
	case "switch":
		{
			var (
				name = flag.Arg(1)
				from = flag.Arg(2)
				to   = flag.Arg(3)
			)
			switches(ctx, svc, name, from, to)
		}
	case "similar":
		{
			var name = flag.Arg(1)
			similar(ctx, svc, name)
		}
	default:
		log.Fatal("unknown command")
	}
}

func dhcp(ctx context.Context, svc logserver.Service, mac uint64, from, to string) {
	logs, err := svc.GetDHCPLogs(ctx, mac, from, to)
	if err != nil {
		log.Fatalf("error getting DHCP logs of %d: %s", mac, err)
	}

	for _, l := range logs.Logs {
		fmt.Printf("DHCP logs for %d:\nIP:%s, Time: %s\nMessage: %s", mac, l.IP, l.TimeStamp, l.Message)
	}
}

func switches(ctx context.Context, svc logserver.Service, name, from, to string) {
	logs, err := svc.GetSwitchLogs(ctx, name, from, to)
	if err != nil {
		log.Fatalf("error getting switch logs of %s: %s", name, err)
	}

	for _, l := range logs.Logs {
		fmt.Printf("Logs from %s switch:\nIP:%s, Time: %s\nMessage: %s", name, l.IP, l.TimeStamp, l.Message)
	}
}

func similar(ctx context.Context, svc logserver.Service, name string) {
	names, err := svc.GetSimilarSwitches(ctx, name)
	if err != nil {
		log.Fatalf("error getting similar to %s switches: %s", name, err)
	}

	for _, s := range names.Sws {
		fmt.Printf("%s: %s\n", s.Name, s.IP)
	}
}
