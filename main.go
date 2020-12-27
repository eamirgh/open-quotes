package main

import (
	"log"
	"net/http"
	"time"

	"github.com/eamirgh/open-quotes/conf"
	"github.com/eamirgh/open-quotes/handler"
	"github.com/gorilla/mux"
)

func main() {
	conf.Init()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index).Methods("GET")
	r.HandleFunc("/ping", handler.Ping).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
