package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerIndex struct{}

func (p *ControllerIndex) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)

	newCtx.Render("index.tpl",
		map[string]interface{}{
			"content": "admin login",
		})
}
