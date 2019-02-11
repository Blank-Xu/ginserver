package routers

import (
	"github.com/gin-contrib/sessions"

	"ginserver/controllers/admin"
	"ginserver/controllers/admin/user_set"
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
	groupAdmin.Use(admin.AuthSession("roleId", enforcer, "/admin/login"))
	// admin root router
	groupAdmin.GET("/", new(admin.ControllerIndex).Get)
	groupAdmin.GET("about", new(admin.ControllerAbout).Get)

	groupUserSet := groupAdmin.Group("user_set")
	ctlInfo := new(user_set.ControllerInfo)
	groupUserSet.GET("info", ctlInfo.Get)
	groupUserSet.POST("info", ctlInfo.Post)
	ctlChangePwd := new(user_set.ControllerChangePwd)
	groupUserSet.GET("change_pwd", ctlChangePwd.Get)
	groupUserSet.POST("change_pwd", ctlChangePwd.Post)
}
