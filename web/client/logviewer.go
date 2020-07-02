package main

import (
	"flag"
	"log"

	"git.sgu.ru/ultramarine/logserver/web/client/handlers"
	"git.sgu.ru/ultramarine/logserver/web/client/server"

	"github.com/gorilla/mux"
)

var confpath = flag.String("conf", "logviewer.conf.toml", "")

func main() {
	flag.Parse()

	err := server.Init(*confpath)
	if err != nil {
		log.Fatalf("Can't init programm: %v", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/login", handlers.LoginHandler)

	router.HandleFunc("/get/dhcp", handlers.GetDHCPLogsHandler).Methods("GET")
	router.HandleFunc("/get/switch", handlers.GetSwitchLogsHandler).Methods("GET")
	router.HandleFunc("/get/similar", handlers.GetSimilarSwitchesHandler).Methods("GET")
}
