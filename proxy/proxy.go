package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerProxyRequest(c *gin.Context) {
	response, err := http.DefaultTransport.RoundTrip(c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	c.Status(response.StatusCode)
	for key, value := range response.Header {
		c.Writer.Header()[key] = value
	}

	_, err = io.Copy(c.Writer, response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
