package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAdmin struct {
	Context
}

func (p *ControllerAdmin) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx, false) {
		return
	}
	p.Render("admin.tpl",
		map[string]interface{}{
			"content": "home",
		})
}
