package casbin

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var Default *casbin.Enforcer

func Init(file string) error {
	Default, err := casbin.NewEnforcer(file, false)
	if err != nil {
		return fmt.Errorf("create casbin enforcer failed, err: " + err.Error())
	}

	// load rules
	if err := addAllPermissionForUser(Default); err != nil {
		return fmt.Errorf("load casbin role menu policy failed, err: %v", err)
	}
	if err := addRoleForUser(Default); err != nil {
		return fmt.Errorf("load casbin user role policy failed, err: %v", err)
	}

	return nil
}

func Enforce(context *gin.Context, userId int) (bool, error) {
	return Default.Enforce(strconv.Itoa(userId), context.Request.URL.Path, context.Request.Method)
}
