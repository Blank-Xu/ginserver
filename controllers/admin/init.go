package admin

import (
	"net/http"

	"github.com/casbin/casbin"

	"ginserver/controllers/admin/admins"
	"ginserver/modules/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	cookieName = "ginserver"
)

func Init(r *gin.Engine, enforcer *casbin.Enforcer) {
	var cfg = config.GetConfig()
	if len(cfg.AppName) == 0 {
		cookieName = cfg.AppName
	}

	adminRouter := r.Group("admin")

	// need login
	adminRouter.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/admin/login")
	})

	// use session middleware
	adminRouter.Use(sessions.Sessions(cookieName, newSessionStore()))
	// register login router
	new(Login).RegisterRouter(adminRouter)

	// casbin role check
	adminRouter.Use(authSession(enforcer))

	new(admins.Admins).RegisterRouter(adminRouter)
}
