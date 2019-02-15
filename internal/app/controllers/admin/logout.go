package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Context
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx, false) {
		return
	}
	p.sessionDestroy()
	p.RespRedirect302("/admin/login")
}
