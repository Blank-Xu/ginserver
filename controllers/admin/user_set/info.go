package user_set

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers/admin"
)

type ControllerInfo struct{}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	newCtx := admin.NewContext(ctx)
	newCtx.Render("info.tpl", gin.H{})
}

func (p *ControllerInfo) Post(ctx *gin.Context) {

}
