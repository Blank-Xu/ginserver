package middlewares

import (
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"ginserver/pkg/context"
)

func CasbinEnforce(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.New(c)

		ok, err := enforcer.Enforce(strconv.Itoa(c.GetInt(KeyUserId)), c.Request.URL.Path, c.Request.Method)
		if err != nil {
			ctx.AbortResponseInternalServerError(err)
			return
		}

		if !ok {
			ctx.AbortResponseForbidden()
			return
		}
	}
}
