package admin

import (
	"fmt"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"ginserver/modules/config"
	"ginserver/modules/e"
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
		Domain:   "",
		MaxAge:   cfgSession.MaxAge,
		Secure:   true,
		HttpOnly: true,
	})
	return
}

func authSession(enforcer *casbin.Enforcer, location string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCtx := NewContext(ctx)
		var err error
		if err = newCtx.ParseSession(); err == nil {
			if newCtx.IsLogin() {
				var ok bool
				if ok, err = enforcer.EnforceSafe(newCtx.GetRole(), ctx.Request.URL.Path, ctx.Request.Method); ok {
					ctx.Next()
					return
				}
			}
			e.RespRedirect308(ctx, location)
			return
		}
		logrus.Error(err)
		e.RespErrForbidden(ctx)
	}
}
