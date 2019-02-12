package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerAdmin struct{}

func (p *ControllerAdmin) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)

	newCtx.Render("admin.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
