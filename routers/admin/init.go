package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"ginserver/controllers/admin"
	"ginserver/modules/config"
)

func Init(r *gin.Engine) {
	var cfg = config.GetConfig()
	adminRouter := r.Group("admin")
	// need login
	adminRouter.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/admin/login")
	})

	adminRouter.GET("login", admin.GetLogin)
	adminRouter.POST("login", admin.PostLogin)

	// use session middleware
	adminRouter.Use(sessions.Sessions(cfg.AppName, newSessionStore()))

	registerAdmin(adminRouter)
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
