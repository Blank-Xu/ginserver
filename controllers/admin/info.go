package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerInfo struct{}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	newCtx := NewContext(ctx)
	newCtx.Render("info.tpl", gin.H{})
}

func (p *ControllerInfo) Post(ctx *gin.Context) {

}
