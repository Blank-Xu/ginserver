package casbin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Enforce(context *gin.Context, userId int) bool {
	ok, err := enforcer.EnforceSafe(strconv.Itoa(userId), context.Request.URL.Path, context.Request.Method)
	if err != nil {
		logrus.Error(err)
	}
	return ok
}
