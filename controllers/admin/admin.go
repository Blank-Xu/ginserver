package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAdmin struct{}

func (p *ControllerAdmin) Get(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}

	c.Render("admin.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
