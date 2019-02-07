package admin

import (
	"github.com/gin-gonic/gin"
)

type index struct{}

func (p *index) registerRouter(r *gin.RouterGroup) {
	r.GET("/", p.Get)
}

func (p *index) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)
	if !newCtx.SessionParse() {
		return
	}
	newCtx.Render("admin.html", map[string]interface{}{
		"content": "admin login",
	})
}
