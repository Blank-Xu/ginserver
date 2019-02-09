package admin

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthSession(sessionKey string, enforcer *casbin.Enforcer, location string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if value := sessions.Default(ctx).Get(sessionKey); value != nil {
			if roleId, ok := value.(int); ok {
				if ok, _ = enforcer.EnforceSafe(roleId, ctx.Request.URL.Path, ctx.Request.Method); ok {
					ctx.Next()
					return
				}
			}
		}
		ctx.Redirect(http.StatusFound, location)
		ctx.Abort()
	}
}
