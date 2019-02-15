package controllers

import (
	"net/http"

	"ginserver/init/config"

	"github.com/gin-gonic/gin"
)

type ControllerIndex struct{}

func (p *ControllerIndex) Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
		})
}
