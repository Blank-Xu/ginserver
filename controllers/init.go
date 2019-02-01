package controllers

import (
	"net/http"
	"time"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/e"
	"ginserver/modules/middleware"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const assetsFile = "./assets"

var (
	router = gin.New()
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
	// cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusMethodNotAllowed))
	})

	// load casbin
	casbinEnforcer := casbin.NewEnforcer(cfg.RbacFile, &models.SCasbin{})
	// register routers
	registerRouter(router, casbinEnforcer)
}

func GetRouter() *gin.Engine {
	return router
}
