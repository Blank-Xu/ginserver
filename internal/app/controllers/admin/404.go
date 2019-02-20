package admin

import (
	"github.com/gin-gonic/gin"
)

type Controller404 struct {
	Context
}

func (p *Controller404) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx, false) {
		return
	}
	p.Render("admin404.tpl", nil)
}
