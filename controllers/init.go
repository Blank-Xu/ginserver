package controllers

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/e"
	"ginserver/modules/middleware"
)

const assetsFile = "./assets"

var (
	router         = gin.New()
	casbinEnforcer *casbin.Enforcer
)

func Init() {
	var cfg = config.GetConfig()

	// set global setting
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
	}
	gin.SetMode(cfg.RunMode)

	// set global middleware
	router.Use(middleware.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())

	router.Static(cfg.AssetsFile, assetsFile)
	router.LoadHTMLGlob(cfg.ViewFile + "/*")

	router.NoRoute(func(c *gin.Context) {
		// if utils.IsAjax(c.Request.Header) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
		// 	return
		// }
		// c.HTML(http.StatusNotFound, "404.html", nil)
	})
	router.NoMethod(func(c *gin.Context) {
		// if utils.IsAjax(c.Request.Header) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusMethodNotAllowed))
		// 	return
		// }
		// c.HTML(http.StatusMethodNotAllowed, "405.html", nil)
	})

	// load casbin
	casbinEnforcer = casbin.NewEnforcer(cfg.RbacFile, &models.SCasbin{})

	// register routers
	registerRouter(router)
}

func GetRouter() *gin.Engine {
	return router
}

func GetCasbin() *casbin.Enforcer {
	return casbinEnforcer
}
