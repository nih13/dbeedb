package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a route for the root path
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")

	})

	// Start the server
	r.Run(":8080")
}
