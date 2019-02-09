package routers

import (
	"ginserver/controllers/api"
	"ginserver/controllers/api/v1/admins"
)

func registerApiRouter() {
	groupApi := router.Group("api")
	groupApi.GET("/", api.Index)
	// jwt and casbin auth
	groupApi.Use(api.AuthJwt(enforcer))

	groupV1 := groupApi.Group("v1")

	ctlAdmins := new(admins.ControllerAdmins)
	groupV1.GET("admins/:id", ctlAdmins.GetOne)
	groupV1.GET("admins", ctlAdmins.Get)
	groupV1.POST("admins", ctlAdmins.Post)
	groupV1.PUT("admins/:id", ctlAdmins.Put)
	groupV1.DELETE("admins/:id", ctlAdmins.Delete)
}
