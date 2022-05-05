package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serverRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test")
}

func main() {
	router := gin.Default()
	router.GET("/test", serverRequest)
	router.Run("localhost:8080")
}
