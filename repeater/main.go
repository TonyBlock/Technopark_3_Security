package main

import (
	"Technopark_3_Security/repositories"
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/cavaliergopher/grab/v3"

	"github.com/gorilla/mux"
)

var paramsFilePath = "./params"

func main() {
	if _, err := os.Stat(paramsFilePath); errors.Is(err, os.ErrNotExist) {
		_, err := grab.Get(paramsFilePath, "https://raw.githubusercontent.com/PortSwigger/param-miner/master/resources/params")
		if err != nil {
			log.Fatal(err)
		}
	}

	requestStore, err := repositories.CreatePostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	repeater := Repeater{Store: requestStore}
	router := mux.NewRouter()
	{
		router.HandleFunc("/requests", repeater.HandleGetRequests)
		router.HandleFunc("/requests/{id:[0-9]+}", repeater.HandleRepeatRequest)
		router.HandleFunc("/requests/{id:[0-9]+}/scan", repeater.HandleScanRequest)
	}

	server := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(server.ListenAndServe())
}
