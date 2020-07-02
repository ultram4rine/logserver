package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"git.sgu.ru/ultramarine/logserver"
	"git.sgu.ru/ultramarine/logserver/client"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Conf struct {
	CertPath      string `toml:"cert"`
	GRPCServer    string `toml:"grpc_server"`
	SessionKey    string `toml:"session_key"`
	EncryptionKey string `toml:"encryption_key"`
	LDAP          ldap   `toml:"ldap"`
}

var Core struct {
	GRPC  logserver.Service
	Store *sessions.CookieStore
}

func Init(confpath string) (err error) {
	if _, err := toml.DecodeFile(confpath, &Conf); err != nil {
		return fmt.Errorf("error decoding config file from %s", confpath)
	}

	if Conf.SessionKey == "" {
		return errors.New("Empty session key")
	}
	if Conf.EncryptionKey == "" {
		return errors.New("Empty encryption key")
	}

	Core.Store = sessions.NewCookieStore([]byte(Conf.SessionKey), []byte(Conf.EncryptionKey))

	b, err := ioutil.ReadFile(Conf.CertPath)
	if err != nil {
		return fmt.Errorf("error reading certificate authority from %s: %s", Conf.CertPath, err)
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return errors.New("failed to append certificates")
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}

	conn, err := grpc.Dial(Conf.GRPCServer, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), grpc.WithTimeout(1*time.Second))
	if err != nil {
		return fmt.Errorf("cannot connect to %s: %s", Conf.GRPCServer, err)
	}
	defer conn.Close()

	Core.GRPC = client.New(conn)

	return nil
}

type ldap struct {
	Host     string `envconfig:"host"`
	BindDN   string `envconfig:"bind_dn"`
	BindPass string `envconfig:"bind_pass"`
	BaseDN   string `envconfig:"base_dn"`
}
