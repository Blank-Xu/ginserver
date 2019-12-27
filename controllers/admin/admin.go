package admin

import (
	"ginserver/global"

	"github.com/gin-gonic/gin"
)

type ControllerAdmin struct {
	Controller
}

func (p *ControllerAdmin) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("admin.tpl",
		map[string]interface{}{
			"content": global.AppName + global.Version,
		})
}
