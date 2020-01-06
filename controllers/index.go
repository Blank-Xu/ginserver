package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/global"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"AppName": global.AppName,
		"Version": global.Version,
	})
}
