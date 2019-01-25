package router

import (
	"fmt"

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

var router = gin.New()

func Init() {
	var cfg = config.GetConfig()

	// set gin's global setting
	if cfg.RunMode != gin.DebugMode {
		gin.DisableConsoleColor()
	}
	gin.SetMode(cfg.RunMode)

	// set gin's global middleware
	router.Use(middleware.Logger(log.GetLog(), assetsFile))
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions(cfg.AppName, newSessionStore()))

	router.Static(cfg.AssetsFile, assetsFile)
	router.LoadHTMLGlob(cfg.ViewFile + "/*")

	e := casbin.NewEnforcer("config/rbac_model.ini", &model.Casbin{})

	// register routers
	registerRouter()
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

func GetRouter() *gin.Engine {
	return router
}

func registerRouter() {
	index()
}
