package routers

import (
	"net/http"

	"ginserver/init/config"
	"ginserver/internal/app/controllers/admin"
	"ginserver/tools/casbin"
	"ginserver/tools/func_map"
	"ginserver/tools/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cookieName = "ginserver"

func registerAdminRouter(router *gin.Engine) {
	var cfg = config.GetConfig()
	if len(cfg.AppName) > 0 {
		cookieName = cfg.AppName
	}

	// statics and templates
	router.Static("/static", cfg.StaticDir)
	router.StaticFS("/assets", http.Dir("./assets"))
	router.StaticFile("/favicon.ico", cfg.StaticDir+"/favicon.ico")
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
		groupAdmin.GET("logout", middleware.SessionDestroy(), new(admin.ControllerLogout).Get)

		groupAdmin.Use(middleware.SessionAuth("/admin/login"))
		groupAdmin.GET("/", new(admin.ControllerAdmin).Get)
		groupAdmin.GET("404", new(admin.Controller404).Get)
		groupAdmin.GET("500", new(admin.Controller500).Get)
		groupAdmin.GET("about", new(admin.ControllerAbout).Get)
		groupAdmin.GET("info", new(admin.ControllerInfo).Get)
		groupAdmin.POST("info", new(admin.ControllerInfo).Post)
		groupAdmin.GET("change_pwd", new(admin.ControllerChangePwd).Get)
		groupAdmin.POST("change_pwd", new(admin.ControllerChangePwd).Post)

		groupAdmin.Use(middleware.CasbinEnforce(casbin.GetEnforcer()))
	}
}
