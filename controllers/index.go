package controllers

import (
	"net/http"

	"ginserver/modules/config"

	"github.com/gin-gonic/gin"
)

type Index struct{}

func (p *Index) RegisterRouter(r *gin.Engine) {
	r.GET("/", p.Get)
}

func (p *Index) Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
		})
}
