package middleware

import (
	"net/http"
	"strconv"

	"ginserver/tools/e"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func CasbinEnforce(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enforcer.Enforce(strconv.Itoa(c.GetInt(KeyUserId)), c.Request.URL.Path, c.Request.Method) {
			c.AbortWithStatusJSON(e.RespErrHttp(http.StatusForbidden))
			return
		}
	}
}
