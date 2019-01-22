package router

import (
	"ginserver/module/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index() {
	router.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, config.GetConfig().Server.AppName)
	})
}
