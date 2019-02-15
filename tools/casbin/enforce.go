package casbin

import (
	"strconv"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

var defaultEnforcer *casbin.Enforcer

func GetEnforcer() *casbin.Enforcer {
	return defaultEnforcer
}

func SetEnforcer(enforcer *casbin.Enforcer) {
	defaultEnforcer = enforcer
}

func EnforceSafe(context *gin.Context, userId int) (bool, error) {
	return defaultEnforcer.EnforceSafe(strconv.Itoa(userId), context.Request.URL.Path, context.Request.Method)
}

func Enforce(context *gin.Context, userId int) bool {
	return defaultEnforcer.Enforce(strconv.Itoa(userId), context.Request.URL.Path, context.Request.Method)
}
