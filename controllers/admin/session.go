package admin

import (
	"fmt"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"ginserver/modules/config"
)

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

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   cfgSession.MaxAge,
		HttpOnly: true,
	})
	return
}

func authSession(enforcer *casbin.Enforcer, location string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCtx := NewContext(ctx)
		if newCtx.SessionParse() {
			if ok, _ := enforcer.EnforceSafe(newCtx.GetRoleId(), newCtx.Request.URL.Path, newCtx.Request.Method); ok {
				newCtx.Next()
				return
			}
		}
		newCtx.RespRedirect302(location)
		newCtx.Abort()
	}
}
