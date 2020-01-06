package casbin

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var defaultEnforcer *casbin.Enforcer

func Init(file string) error {
	enforcer, err := casbin.NewEnforcer(file, false)
	if err != nil {
		return fmt.Errorf("create casbin enforcer failed, err: " + err.Error())
	}

	// load rules
	if err := addAllPermissionForUser(enforcer); err != nil {
		return fmt.Errorf("load casbin role menu policy failed, err: %v", err)
	}
	if err := addRoleForUser(enforcer); err != nil {
		return fmt.Errorf("load casbin user role policy failed, err: %v", err)
	}

	defaultEnforcer = enforcer

	return nil
}

func GetEnforcer() *casbin.Enforcer {
	return defaultEnforcer
}

func Enforce(context *gin.Context, userId int) (bool, error) {
	return defaultEnforcer.Enforce(strconv.Itoa(userId), context.Request.URL.Path, context.Request.Method)
}
