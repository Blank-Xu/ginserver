package api

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func AuthJwt(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
