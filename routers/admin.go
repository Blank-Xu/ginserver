package routers

import (
	"github.com/gin-contrib/sessions"

	"ginserver/controllers/admin"
	"ginserver/modules/config"
)

var cookieName = "ginserver"

func registerAdminRouter() {
	var cfg = config.GetConfig()
	if len(cfg.AppName) > 0 {
		cookieName = cfg.AppName
	}

	groupAdmin := router.Group("admin")
	// use session middleware
	groupAdmin.Use(sessions.Sessions(cookieName, newSessionStore()))
	// register login router
	login := new(admin.ControllerLogin)
	groupAdmin.GET("login", login.Get)
	groupAdmin.POST("login", login.Post)
	// register logout router
	groupAdmin.GET("logout", new(admin.ControllerLogout).Get)
	// casbin roleId check
	groupAdmin.Use(admin.AuthSession(enforcer, "/admin/login"))
	// admin root router
	groupAdmin.GET("/", new(admin.ControllerIndex).Get)
}
