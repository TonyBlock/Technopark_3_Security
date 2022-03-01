package main

import (
	"Technopark_3_Security/models"
	"Technopark_3_Security/proxyHandlers"
	"Technopark_3_Security/repositories"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

var proxyAddress = ":8080"

func main() {
	requestStore, err := repositories.CreatePostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr: proxyAddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var request models.Request
			if err = request.SetHTTPRequest(*req); err != nil {
				log.Println(err)
			}
			if err = requestStore.Set(request); err != nil {
				log.Println(err)
			}

			if req.Method == http.MethodConnect {
				proxyHandlers.HandleConnectHTTPS(w, req)
			} else {
				proxyHandlers.HandleHTTP(w, req)
			}
		}),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	fmt.Println(server.ListenAndServe())
}
