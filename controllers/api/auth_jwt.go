package api

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func authJwt(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
