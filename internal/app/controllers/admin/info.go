package admin

import (
	"ginserver/internal/app/models"

	"github.com/gin-gonic/gin"
)

type ControllerInfo struct {
	Controller
}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("info.tpl", gin.H{})
}

func (p *ControllerInfo) Post(ctx *gin.Context) {
	p.New(ctx)
	var err error
	req := new(models.SUserInfoUpdate)
	if err = p.ShouldBind(req); err != nil {
		p.RespErrInvalidParams()
		return
	}

}
