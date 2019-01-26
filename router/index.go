package router

import (
	"net/http"

	"ginserver/module/config"

	"github.com/gin-gonic/gin"
)

func registerIndex(r *gin.Engine) {
	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
		})
	})
}
