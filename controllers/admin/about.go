package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAbout struct {
	Context
}

func (p *ControllerAbout) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx) {
		return
	}
	p.Render("about.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
