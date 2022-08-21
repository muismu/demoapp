package main

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/muismu/demoapp/handlers"
)

//go:embed templates
var Files embed.FS

func main() {
	router := gin.Default()
	tmpl, err := template.ParseFS(Files, "templates/*")
	if err != nil {
		panic(err.Error())
	}
	router.SetHTMLTemplate(tmpl)
	router.GET("/ping", handlers.Ping)
	v1 := router.Group("/v1")
	v1.Use(gin.BasicAuth(gin.Accounts{"mu": "mu"}))
	{
		v1.GET("/user", handlers.ShowUserAuthInfo)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/ping", handlers.Ping)
	}
	router.Run(":80")
}
