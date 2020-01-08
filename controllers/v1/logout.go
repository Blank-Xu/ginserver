package v1

import (
	"github.com/gin-gonic/gin"

	"ginserver/pkg/context"
	"ginserver/pkg/middlewares"
)

func LogoutGet(c *gin.Context) {
	ctx := context.New(c)

	if err := middlewares.SessionDestroy(c); err != nil {
		ctx.AbortResponseInternalServerError(err)
		return
	}

	ctx.Redirect302("/")
}
