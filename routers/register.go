package routers

import (
	"net/http"
	"strings"
	"time"

	"ginserver/controllers"
	"ginserver/global"
	"ginserver/pkg/e"
	"ginserver/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var defaultRouter = gin.New()

func GetDefaultRouter() *gin.Engine {
	return defaultRouter
}

func Register() {
	register(defaultRouter)
}

func register(router *gin.Engine) {
	// set global setting
	gin.SetMode(global.RunMode)
	if global.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
		gin.DefaultErrorWriter = logrus.StandardLogger().Out
		router.Use(middlewares.Logger(logrus.StandardLogger()))
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
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
		} else {
			c.Redirect(http.StatusFound, "/admin/404")
		}
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(e.RespErrHttp(http.StatusMethodNotAllowed))
	})
	// home page
	router.GET("/", new(controllers.ControllerIndex).Get)

	registerAdminRouter(router)
	registerApiRouter(router)
}
