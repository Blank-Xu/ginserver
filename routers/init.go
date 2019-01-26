package routers

import (
	"net/http"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/errors"
	"ginserver/modules/log"
	"ginserver/modules/middleware"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
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
	router.Use(middleware.Logger(log.GetLog()))
	router.Use(gin.Recovery())

	router.Static(cfg.AssetsFile, assetsFile)
	router.LoadHTMLGlob(cfg.ViewFile + "/*")

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(errors.RespHttpError(http.StatusNotFound))
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(errors.RespHttpError(http.StatusMethodNotAllowed))
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
