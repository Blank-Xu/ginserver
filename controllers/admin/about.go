package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAbout struct {
	Controller
}

func (p *ControllerAbout) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("about.tpl",
		map[string]interface{}{
			"content": "about",
		})
}
