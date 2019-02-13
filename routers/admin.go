package routers

import (
	"github.com/gin-contrib/sessions"

	"ginserver/controllers/admin"
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
		groupAdmin.GET("/", new(admin.ControllerAdmin).Get)
		groupAdmin.GET("about", new(admin.ControllerAbout).Get)
		groupAdmin.GET("info", new(admin.ControllerInfo).Get)
		groupAdmin.POST("info", new(admin.ControllerInfo).Post)
		groupAdmin.GET("change_pwd", new(admin.ControllerChangePwd).Get)
		groupAdmin.POST("change_pwd", new(admin.ControllerChangePwd).Post)
	}
}
