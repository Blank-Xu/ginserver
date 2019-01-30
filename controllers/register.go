package controllers

import (
	"ginserver/controllers/admin"
	"ginserver/controllers/api"

	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {
	new(Index).RegisterRouter(r)

	admin.Init(r)

	api.Init(r)
}
