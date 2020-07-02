package main

import (
	"flag"
	"log"
	"net/http"

	"git.sgu.ru/ultramarine/logserver/web/client/handlers"
	"git.sgu.ru/ultramarine/logserver/web/client/server"

	"github.com/gorilla/mux"
)

var confpath = flag.String("conf", "logviewer.conf.toml", "")

func main() {
	flag.Parse()

	err := server.Init(*confpath)
	if err != nil {
		log.Fatalf("failed to init logviewer: %v", err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static/"))))

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/login", handlers.LoginHandler)

	router.HandleFunc("/get/dhcp", handlers.GetDHCPLogsHandler).Methods("GET")
	router.HandleFunc("/get/switch", handlers.GetSwitchLogsHandler).Methods("GET")
	router.HandleFunc("/get/similar", handlers.GetSimilarSwitchesHandler).Methods("GET")

	err = http.ListenAndServe(":"+server.Conf.App.ListenPort, router)
	if err != nil {
		log.Fatalf("failed to start logviewer: %v", err)
	}
}
