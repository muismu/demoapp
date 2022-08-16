package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "Ok"})
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	v1 := router.Group("/v1")
	v1.Use(gin.BasicAuth(gin.Accounts{"mu": "mu"}))
	{
		v1.GET("/ping", ping)
	}
	router.Run(":80")
}
