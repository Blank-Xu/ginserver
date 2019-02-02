package controllers

import (
	"ginserver/controllers/admin"
	"ginserver/controllers/api"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func registerRouter(r *gin.Engine, enforcer *casbin.Enforcer) {
	// register swagger doc router
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// home page
	new(index).registerRouter(r)

	admin.Init(r, enforcer)

	api.Init(r, enforcer)
}
