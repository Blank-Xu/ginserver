package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerLogout struct{}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}
	c.SessionDestroy()
	c.RespRedirect302("/admin/login")
}
