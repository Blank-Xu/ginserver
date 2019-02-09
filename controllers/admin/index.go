package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerIndex struct{}

func (p *ControllerIndex) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)
	if !newCtx.SessionParse() {
		return
	}
	newCtx.Render("admin.html", map[string]interface{}{
		"content": "admin login",
	})
}
