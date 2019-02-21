package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Controller
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	p.New(ctx)
	p.RespRedirect302("/admin/login")
}
