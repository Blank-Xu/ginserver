package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerLogout struct{}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)
	newCtx.SessionDestroy()
	newCtx.RespRedirect302("/admin/login")
}
