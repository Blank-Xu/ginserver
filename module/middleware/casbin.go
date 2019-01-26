package middleware

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Casbin(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
