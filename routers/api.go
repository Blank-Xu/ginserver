package routers

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"ginserver/controllers/api"
	"ginserver/controllers/api/v1/admins"
)

func registerApiRouter() {
	// register swagger doc router
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	groupApi := router.Group("api")
	groupApi.GET("/", api.Index)
	// jwt and casbin auth
	groupApi.Use(api.AuthJwt(enforcer))

	apiV1 := groupApi.Group("v1")
	{
		apiV1.GET("admins/:id", new(admins.ControllerAdmins).GetOne)
		apiV1.GET("admins", new(admins.ControllerAdmins).Get)
		apiV1.POST("admins", new(admins.ControllerAdmins).Post)
		apiV1.PUT("admins/:id", new(admins.ControllerAdmins).Put)
		apiV1.DELETE("admins/:id", new(admins.ControllerAdmins).Delete)
	}
}
