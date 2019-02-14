package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Context
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx) {
		return
	}
	p.SessionDestroy()
	p.RespRedirect302("/admin/login")
}
