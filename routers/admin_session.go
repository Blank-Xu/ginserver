package routers

import (
	"github.com/gin-contrib/sessions"
)

func newSessionStore() (store sessions.Store) {
	// var (
	// 	cfgSession = config.GetConfig().Session
	// 	cfgRedis   = config.GetConfig().Redis
	// 	err        error
	// )
	//
	// switch cfgSession.Provider {
	// case "redis":
	// 	store, err = redis.NewStore(30, "tcp", cfgRedis.Host+":"+cfgRedis.Port, cfgRedis.Password, []byte(cfgSession.Secret))
	// 	if err != nil {
	// 		panic(fmt.Sprintf("create redis session err: [%v]", err))
	// 	}
	// case "memstore":
	// 	store = memstore.NewStore([]byte(cfgSession.Secret))
	// default:
	// 	panic("load session config error")
	// }
	//
	// store.Options(sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   cfgSession.MaxAge,
	// 	HttpOnly: true,
	// })
	return
}
