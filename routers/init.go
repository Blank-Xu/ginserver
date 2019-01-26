package routers

import (
	"fmt"
	"net/http"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/errors"
	"ginserver/modules/log"
	"ginserver/modules/middleware"

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
		c.AbortWithStatusJSON(errors.ErrorHttpCodeJson(http.StatusNotFound))
	})
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(errors.ErrorHttpCodeJson(http.StatusMethodNotAllowed))
	})

	router.Use(sessions.Sessions(cfg.AppName, newSessionStore()))

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
