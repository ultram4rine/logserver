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

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var server = flag.String("logserver", "localhost:8908", "")

func main() {
	flag.Parse()

	ctx := context.Background()

	conn, err := grpc.Dial(*server, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalf("cannot connect to %s: %s", *server, err)
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
		}
	case "similar":
		{
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
	fmt.Println(logs)
}
