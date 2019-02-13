package admin

import (
	"github.com/gin-gonic/gin"
)

type ControllerInfo struct{}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}
	c.Render("info.tpl", gin.H{})
}

func (p *ControllerInfo) Post(ctx *gin.Context) {

}
