package router

import (
	"ginserver/module/config"
	"ginserver/module/log"
	"ginserver/module/middleware"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func Init() {
	var cfg = config.GetConfig()

	// set gin's global setting
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
	}
	gin.SetMode(cfg.RunMode)

	// set gin's middleware
	router.Use(middleware.Logger(log.GetLog()), gin.Recovery())

	// register routers
	registerRouter()
}

func GetRouter() *gin.Engine {
	return router
}

func registerRouter() {
	index()
}
