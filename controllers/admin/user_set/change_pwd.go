package user_set

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers/admin"
)

type ControllerChangePwd struct{}

func (p *ControllerChangePwd) Get(ctx *gin.Context) {
	newCtx := admin.NewContext(ctx)
	newCtx.Render("change_pwd.tpl", gin.H{})
}

func (p *ControllerChangePwd) Post(ctx *gin.Context) {

}
