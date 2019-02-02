package controllers

import (
	"net/http"

	"ginserver/modules/config"

	"github.com/gin-gonic/gin"
)

type index struct{}

func (p *index) registerRouter(r *gin.Engine) {
	r.GET("/", p.get)
}

func (p *index) get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
		})
}
