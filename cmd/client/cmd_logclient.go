package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
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

	b, err := ioutil.ReadFile(conf.Cert)
	if err != nil {
		log.Fatalf("error reading certificate authority from %s: %s", conf.Cert, err)
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatal("failed to append certificates")
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}

	conn, err := grpc.Dial(conf.Server, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), grpc.WithTimeout(1*time.Second))
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

			if mac == "" {
				log.Fatal("please provide a MAC address which logs you want to see")
			}

			dhcp(ctx, svc, mac, from, to)
		}
	case "switch":
		{
			var (
				name = flag.Arg(1)
				from = flag.Arg(2)
				to   = flag.Arg(3)
			)

			if name == "" {
				log.Fatal("please provide switch name which logs you want to see")
			}

			switches(ctx, svc, name, from, to)
		}
	case "similar":
		{
			var name = flag.Arg(1)

			if name == "" {
				log.Fatal("please provide similar name")
			}

			similar(ctx, svc, name)
		}
	default:
		log.Fatal("unknown command")
	}
}

func dhcp(ctx context.Context, svc logserver.Service, mac, from, to string) {
	logs, err := svc.GetDHCPLogs(ctx, mac, from, to)
	if err != nil {
		log.Fatalf("error getting DHCP logs of %d: %s", mac, err)
	}

	fmt.Printf("DHCP logs for %s:\n", mac)

	for _, l := range logs.Logs {
		fmt.Printf("IP: %s, Time: %s\nMessage: %s\n\n", l.IP, l.TimeStamp, l.Message)
	}
}

func switches(ctx context.Context, svc logserver.Service, name, from, to string) {
	logs, err := svc.GetSwitchLogs(ctx, name, from, to)
	if err != nil {
		log.Fatalf("error getting switch logs of %s: %s", name, err)
	}

	fmt.Printf("Logs from %s switch:\n", name)

	for _, l := range logs.Logs {
		fmt.Printf("IP: %s, Time: %s\nMessage: %s\n\n", l.IP, l.TimeStamp, l.Message)
	}
}

func similar(ctx context.Context, svc logserver.Service, name string) {
	names, err := svc.GetSimilarSwitches(ctx, name)
	if err != nil {
		log.Fatalf("error getting similar to %s switches: %s", name, err)
	}

	fmt.Printf("Similars to %s:\n", name)

	for _, s := range names.Sws {
		fmt.Printf("%s: %s\n", s.Name, s.IP)
	}
}
