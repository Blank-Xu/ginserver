package admin

import (
	"ginserver/modules/config"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const loginLocation = "login"

var cookieName = "ginserver"

func Init(r *gin.Engine, enforcer *casbin.Enforcer) {
	var cfg = config.GetConfig()
	if len(cfg.AppName) > 0 {
		cookieName = cfg.AppName
	}

	adminRouter := r.Group("admin")

	// use session middleware
	adminRouter.Use(sessions.Sessions(cookieName, newSessionStore()))
	// register login router
	new(login).registerRouter(adminRouter)
	new(logout).registerRouter(adminRouter, loginLocation)
	// casbin role check
	adminRouter.Use(authSession(enforcer, loginLocation))
	// admin root router
	new(index).registerRouter(adminRouter)
	new(admin).registerRouter(adminRouter)
}
