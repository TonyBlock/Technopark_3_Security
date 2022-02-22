package main

import (
	"Technopark_3_Security/proxy"
	"crypto/tls"
	"fmt"
	//"log"
	"net/http"
	//"github.com/gin-gonic/gin"
)

var proxyAddress = ":8080"

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//router := gin.Default()
	//
	//router.NoRoute(proxy.HandlerProxyRequest)

	server := &http.Server{
		Addr:         proxyAddress,
		Handler:      http.HandlerFunc(proxy.HandlerProxyRequest),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	fmt.Println(server.ListenAndServe())
	//
	//err := router.Run(proxyAddress)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
