package casbin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthSession(sessionKey string, location string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if value := sessions.Default(ctx).Get(sessionKey); value != nil {
			if id, ok := value.(int); ok {
				if ok, _ = enforcer.EnforceSafe(id, ctx.Request.URL.Path, ctx.Request.Method); ok {
					ctx.Next()
					return
				}
			}
		}
		ctx.Redirect(http.StatusFound, location)
		ctx.Abort()
	}
}
