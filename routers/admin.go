package routers

import (
	"github.com/gin-contrib/sessions"

	"ginserver/controllers/admin"
	"ginserver/controllers/admin/user_set"
	"ginserver/modules/config"
	"ginserver/modules/func_map"
)

var cookieName = "ginserver"

func registerAdminRouter() {
	var cfg = config.GetConfig()
	if len(cfg.AppName) > 0 {
		cookieName = cfg.AppName
	}

	// statics and templates
	router.Static("/statics", cfg.StaticDir)
	router.SetFuncMap(func_map.GetFunMap())
	router.HTMLRender = loadTemplates(cfg.TemplateDir)

	groupAdmin := router.Group("admin")
	// use session middleware
	groupAdmin.Use(sessions.Sessions(cookieName, newSessionStore()))
	{
		// register login router
		groupAdmin.GET("login", new(admin.ControllerLogin).Get)
		groupAdmin.POST("login", new(admin.ControllerLogin).Post)
		// register logout router
		groupAdmin.GET("logout", new(admin.ControllerLogout).Get)
	}
	// casbin roleId check
	groupAdmin.Use(admin.AuthSession("roleId", enforcer, "/admin/login"))
	{
		// admin root router
		groupAdmin.GET("/", new(admin.ControllerAdmin).Get)
		groupAdmin.GET("about", new(admin.ControllerAbout).Get)
	}
	userSet := groupAdmin.Group("user_set")
	{
		userSet.GET("info", new(user_set.ControllerInfo).Get)
		userSet.POST("info", new(user_set.ControllerInfo).Post)
		userSet.GET("change_pwd", new(user_set.ControllerChangePwd).Get)
		userSet.POST("change_pwd", new(user_set.ControllerChangePwd).Post)
	}
}
