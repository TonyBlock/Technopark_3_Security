package main

import (
	"Technopark_3_Security/repositories"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	requestStore, err := repositories.CreatePostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	repeater := Repeater{Store: requestStore}
	router := mux.NewRouter()
	{
		router.HandleFunc("/requests", repeater.HandleGetRequests)
		router.HandleFunc("/requests/{id:[0-9]+}", repeater.HandleRepeatRequest)
	}

	server := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(server.ListenAndServe())
}
