package routers

import (
	"net/http"

	"ginserver/controllers/admin"
	"ginserver/global"
	"ginserver/pkg/casbin"
	"ginserver/pkg/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func registerAdminRouter(router *gin.Engine) {
	// var cfg = config.GetConfig()
	// if len(cfg.AppName) > 0 {
	// 	cookieName = global.AppName
	// }

	// statics and templates
	// router.Static("/static", cfg.StaticDir)
	router.StaticFS("/assets", http.Dir("./assets"))
	// router.StaticFile("/favicon.ico", cfg.StaticDir+"/favicon.ico")
	// router.SetFuncMap(func_map.GetFunMap())
	// router.HTMLRender = loadTemplates(cfg.TemplateDir)

	groupAdmin := router.Group("admin")
	// use session middleware
	sessionStore, err := global.DefaultConfig.Session.NewStore()
	if err != nil {
		panic(err)
	}
	groupAdmin.Use(sessions.Sessions(global.AppName, sessionStore))
	{
		// register /admin/login router
		groupAdmin.GET("login", new(admin.ControllerLogin).Get)
		groupAdmin.POST("login", new(admin.ControllerLogin).Post)

		groupAdmin.Use(middlewares.SessionAuth("/admin/login"))
		// register /admin/logout router
		groupAdmin.GET("logout", new(admin.ControllerLogout).Get, middlewares.SessionDestroy())
		// register /admin router
		groupAdmin.GET("/", new(admin.ControllerAdmin).Get)
		groupAdmin.GET("404", new(admin.Controller404).Get)
		groupAdmin.GET("500", new(admin.Controller500).Get)
		groupAdmin.GET("about", new(admin.ControllerAbout).Get)
		groupAdmin.GET("info", new(admin.ControllerInfo).Get)
		groupAdmin.POST("info", new(admin.ControllerInfo).Post)
		groupAdmin.GET("change_pwd", new(admin.ControllerChangePwd).Get)
		groupAdmin.POST("change_pwd", new(admin.ControllerChangePwd).Post)

		groupAdmin.Use(middlewares.CasbinEnforce(casbin.GetEnforcer()))
	}
}
