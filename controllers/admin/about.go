package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAbout struct{}

func (p *ControllerAbout) Get(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}

	c.Render("about.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
