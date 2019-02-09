package admin

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func AuthSession(enforcer *casbin.Enforcer, location string) gin.HandlerFunc {
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
