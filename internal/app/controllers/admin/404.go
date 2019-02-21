package admin

import (
	"github.com/gin-gonic/gin"
)

type Controller404 struct {
	Controller
}

func (p *Controller404) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("admin404.tpl", nil)
}
