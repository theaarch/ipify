package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		ip := c.ClientIP()

		c.String(http.StatusOK, ip)
	})

	err := r.Run()
	if err != nil {
		return
	}
}
