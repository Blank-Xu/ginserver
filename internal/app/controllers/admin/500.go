package admin

import (
	"github.com/gin-gonic/gin"
)

type Controller500 struct {
	Context
}

func (p *Controller500) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx, false) {
		return
	}
	p.Render("admin500.tpl", nil)
}
