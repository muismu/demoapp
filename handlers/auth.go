package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserAuthInfo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.HTML(http.StatusOK, "DisplayJson.tmpl", gin.H{"User": user, "Headers": c.Request.Header})
}
