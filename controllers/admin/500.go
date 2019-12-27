package admin

import (
	"github.com/gin-gonic/gin"
)

type Controller500 struct {
	Controller
}

func (p *Controller500) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("admin500.tpl", nil)
}
