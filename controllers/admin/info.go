package admin

import (
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
	if !p.ParseContext(ctx) {
		return
	}

}
