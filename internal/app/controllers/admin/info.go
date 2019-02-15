package admin

import (
	"ginserver/internal/app/models"

	"github.com/gin-gonic/gin"
)

type ControllerInfo struct {
	Context
}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	if !p.ParseContext(ctx) {
		return
	}
	p.Render("info.tpl", gin.H{})
}

func (p *ControllerInfo) Post(ctx *gin.Context) {
	if !p.ParseContext(ctx, false) {
		return
	}
	var err error
	req := new(models.SUserInfoUpdate)
	if err = p.ShouldBind(req); err != nil {
		p.RespErrInvalidParams()
		return
	}

}
