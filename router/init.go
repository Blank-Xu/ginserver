package router

import (
	"fmt"
	"net/http"

	"ginserver/model"
	"ginserver/module/config"
	"ginserver/module/log"
	"ginserver/module/middleware"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
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
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  http.StatusText(http.StatusNotFound),
			})
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed,
			gin.H{
				"code": http.StatusMethodNotAllowed,
				"msg":  http.StatusText(http.StatusMethodNotAllowed),
			})
	})

	router.Use(sessions.Sessions(cfg.AppName, newSessionStore()))

	// load casbin
	casbinEnforcer = casbin.NewEnforcer(cfg.RbacFile, &model.SCasbin{})

	// register routers
	registerRouter(router)
}

func GetRouter() *gin.Engine {
	return router
}

func GetCasbin() *casbin.Enforcer {
	return casbinEnforcer
}

func newSessionStore() (store sessions.Store) {
	var (
		cfgSession = config.GetConfig().Session
		cfgRedis   = config.GetConfig().Redis
		err        error
	)

	switch cfgSession.Provider {
	case "redis":
		store, err = redis.NewStore(30, "tcp", cfgRedis.Host+":"+cfgRedis.Port, cfgRedis.Password, []byte(cfgSession.Secret))
		if err != nil {
			panic(fmt.Sprintf("create redis session err: [%v]", err))
		}
	case "memstore":
		store = memstore.NewStore([]byte(cfgSession.Secret))
	default:
		panic("load session config error")
	}
	return
}
