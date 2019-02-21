package admin

import (
	"github.com/gin-gonic/gin"

	"ginserver/init/config"
)

type ControllerAdmin struct {
	Controller
}

func (p *ControllerAdmin) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("admin.tpl",
		map[string]interface{}{
			"content": config.GetConfig().AppName + config.Version,
		})
}
