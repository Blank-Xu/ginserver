package routers

import (
	"net/http"
	"time"

	"ginserver/controllers"
	"ginserver/modules/config"
	"ginserver/modules/e"
	"ginserver/modules/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var router = gin.New()

func Init() {
	var cfg = config.GetConfig()
	// set global setting
	gin.SetMode(cfg.RunMode)
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
		gin.DefaultErrorWriter = logrus.StandardLogger().Out
		router.Use(middleware.Logger(logrus.StandardLogger()))
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
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusMethodNotAllowed))
	})
	// home page
	router.GET("/", new(controllers.ControllerIndex).Get)

	registerAdminRouter()
	registerApiRouter()
}

func GetRouter() *gin.Engine {
	return router
}
