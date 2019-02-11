package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAbout struct{}

func (p *ControllerAbout) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)

	newCtx.Render("about.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
