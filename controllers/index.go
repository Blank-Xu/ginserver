package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/global"
)

type ControllerIndex struct{}

func (p *ControllerIndex) Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		map[string]string{
			"AppName": global.AppName,
			"Version": global.Version,
		})
}
