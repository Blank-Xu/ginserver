package routers

import (
	"time"

	"github.com/rs/zerolog/log"

	"ginserver/controllers"
	"ginserver/global"
	"ginserver/pkg/context"
	"ginserver/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	defaultRouter := gin.New()

	registerDefault(defaultRouter)

	// registerAdminRouter(router)
	registerApi(defaultRouter)

	return defaultRouter
}

func registerDefault(router *gin.Engine) {
	// set global setting
	gin.SetMode(global.RunMode)
	if global.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
		gin.DefaultErrorWriter = log.Logger
		router.Use(middlewares.Logger(&log.Logger))
	} else {
		pprof.Register(router)
		router.Use(gin.Logger())
	}
	// set global middleware
	router.Use(gin.Recovery())

	// cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Hour * 12,
	}))

	router.NoRoute(func(c *gin.Context) {
		ctx := context.New(c)
		ctx.AbortResponseNotFound()
	})
	router.NoMethod(func(c *gin.Context) {
		ctx := context.New(c)
		ctx.AbortResponseMethodNotAllowed()
	})

	// home index
	router.GET("/", controllers.Index)
	router.POST("/", controllers.Index)
}
