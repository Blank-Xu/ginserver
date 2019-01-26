package controllers

import (
	"net/http"

	"ginserver/modules/config"

	"github.com/gin-gonic/gin"
)

func GetIndex(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			map[string]string{
				"AppName": config.GetConfig().AppName,
				"Version": config.Version,
			})
	})
}
