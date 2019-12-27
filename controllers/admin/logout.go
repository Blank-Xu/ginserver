package admin

import (
	"ginserver/models/log"

	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Controller
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	p.New(ctx)
	p.LogDB(log.TypeLogout, log.LevelInfo)
	p.RespRedirect302("/admin/login")
}
