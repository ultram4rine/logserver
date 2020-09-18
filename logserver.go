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
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"git.sgu.ru/ultramarine/logserver/auth"
	"git.sgu.ru/ultramarine/logserver/conf"
	"git.sgu.ru/ultramarine/logserver/pb"
	"git.sgu.ru/ultramarine/logserver/service"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	if _, err := os.Stat("ui/node_modules"); os.IsNotExist(err) {
		log.Warn("Dependencies of web app are not installed")
		log.Info("Running 'npm install'...")

		cmd := exec.Command("npm", "install")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = "./ui"

		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed install web app dependencies: %s", err)
		}

		log.Info("Dependencies of web app installed")
	}

	cmd := exec.Command("npm", "run", "build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "./ui"

	log.Info("Building web app...")

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to build web app: %s", err)
	}

	log.Info("Web app builded successfully")
}

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

		gwmux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))}
		err = pb.RegisterLogServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%s", conf.Conf.App.ListenPort), opts)
		if err != nil {
			errChan <- err
			return
		}

		spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}

		router := mux.NewRouter()
		router.HandleFunc("/api/auth", auth.Handler)
		router.PathPrefix("/api").Handler(gwmux)
		router.PathPrefix("/").Handler(spa)

		srv := &http.Server{
			Handler:      router,
			Addr:         ":" + conf.Conf.App.GatewayPort,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%v", <-c)
	}()

	log.Info(<-errChan)
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
