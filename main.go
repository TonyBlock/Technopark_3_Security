package main

import (
	"Technopark_3_Security/proxy"
	"fmt"

	"github.com/gin-gonic/gin"
)

var proxyAddress = ":8080"

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(proxy.HandlerProxyRequest)

	err := router.Run(proxyAddress)
	if err != nil {
		fmt.Println(err)
	}
}
