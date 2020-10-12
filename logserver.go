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
	"path/filepath"
	"syscall"
	"time"

	"git.sgu.ru/ultramarine/logserver/auth"
	"git.sgu.ru/ultramarine/logserver/cmd"
	"git.sgu.ru/ultramarine/logserver/conf"
	"git.sgu.ru/ultramarine/logserver/pb"
	"git.sgu.ru/ultramarine/logserver/service"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	confName               = kingpin.Flag("conf", "Path to config file.").Short('c').Default("logserver.conf").String()
	installWEBDependencies = kingpin.Flag("install-spa-dependencies", "Install WEB app dependencies.").Short('i').Bool()
	buildSPA               = kingpin.Flag("build-spa", "Build WEB app.").Short('b').Bool()
	tlsEnabled             = kingpin.Flag("tls", "If used TLS cookies will have Secure flag").Short('s').Bool()
)

func init() {
	kingpin.Parse()

	if *installWEBDependencies {
		log.Info("Running 'npm install'...")
		if err := cmd.InstallWEBDependenciesCmd.Run(); err != nil {
			log.Fatalf("Failed to install web app dependencies: %s", err)
		}
		log.Info("Dependencies of web app installed")
	}

	if *buildSPA {
		log.Info("Building web app...")
		if err := cmd.BuildWEBAppCmd.Run(); err != nil {
			log.Fatalf("Failed to build web app: %s", err)
		}
		log.Info("Web app builded successfully")
	}

	if _, err := os.Stat("ui/node_modules"); os.IsNotExist(err) {
		log.Fatal("Dependencies of web app are not installed.\nRun program with '-i' flag or run 'npm install' in 'ui' folder")
	}
	if _, err := os.Stat("ui/build"); os.IsNotExist(err) {
		log.Fatal("Web app are not built.\nRun program with '-b' flag or run 'npm run build' in 'ui' folder")
	}
}

func main() {
	if err := conf.Load(*confName); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	auth.InitKeysAndCookies(*tlsEnabled)

	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "clickhouse", fmt.Sprintf("%s?username=%s&password=%s&database=%s", conf.Config.DBHost, conf.Config.DBUser, conf.Config.DBPass, conf.Config.DBName))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Info("Connected to ClickHouse database")

	var (
		logger  = log.New()
		svc     = service.LogService{DB: db}
		errChan = make(chan error, 1000)
	)

	go func() {
		listener, err := net.Listen("tcp", ":"+conf.Config.GRPCPort)
		if err != nil {
			errChan <- err
			return
		}

		creds, err := credentials.NewServerTLSFromFile(conf.Config.CertPath, conf.Config.KeyPath)
		if err != nil {
			errChan <- err
			return
		}

		entry := log.NewEntry(logger)
		opts := []grpc_logrus.Option{
			grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
		}
		grpc_logrus.ReplaceGrpcLogger(entry)

		gRPCServer := grpc.NewServer(
			grpc.Creds(creds),
			grpc_middleware.WithUnaryServerChain(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_auth.UnaryServerInterceptor(auth.LDAPAuthFunc),
				grpc_logrus.UnaryServerInterceptor(entry, opts...),
			),
		)
		pb.RegisterLogServiceServer(gRPCServer, svc)

		log.Infof("Started LogServer on %s port", conf.Config.GRPCPort)

		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		b, err := ioutil.ReadFile(conf.Config.ClientCertPath)
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
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
			grpc.WithUnaryInterceptor(grpc_logrus.UnaryClientInterceptor(log.NewEntry(logger))),
		}
		err = pb.RegisterLogServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%s", conf.Config.GRPCPort), opts)
		if err != nil {
			errChan <- err
			return
		}

		spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}

		router := mux.NewRouter()
		router.PathPrefix("/api").Handler(gwmux)
		router.HandleFunc("/auth", auth.Handler)
		router.HandleFunc("/logout", auth.LogoutHandler)
		router.PathPrefix("/").Handler(spa)
		router.Use(auth.TwoCookieAuthMiddleware)

		srv := &http.Server{
			Handler:      router,
			Addr:         ":" + conf.Config.GatewayPort,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		errChan <- srv.ListenAndServe()
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

	if _, err = os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
