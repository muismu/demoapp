package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserAuthInfo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{"User": user})
}
