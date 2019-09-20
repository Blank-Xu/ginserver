package middleware

import (
	"net/http"
	"strconv"

	"ginserver/tools/e"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinEnforce(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, err := enforcer.Enforce(strconv.Itoa(c.GetInt(KeyUserId)), c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.AbortWithStatusJSON(e.RespErrHttp(http.StatusInternalServerError))
			return
		}
		if !ok {
			c.AbortWithStatusJSON(e.RespErrHttp(http.StatusForbidden))
			return
		}
	}
}
