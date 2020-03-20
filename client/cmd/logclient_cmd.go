package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"git.sgu.ru/ultramarine/logserver"
	"git.sgu.ru/ultramarine/logserver/client"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var addr = kingpin.Flag("logserver", "Address of logserver.").Short('s').Default("localhost:8181").String()

func main() {
	kingpin.Parse()

	ctx := context.Background()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalf("cannot connect to %s: %s", *addr, err)
	}
	defer conn.Close()

	svc := client.New(conn)

	args := flag.Args()
	var cmd string

	cmd, args = parse(args)

	switch cmd {
	case "dhcp":
		{
			var mac, from, to string

			mac, args = parse(args)
			from, args = parse(args)
			to, args = parse(args)

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

func parse(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return args[0], args[1:]
}
