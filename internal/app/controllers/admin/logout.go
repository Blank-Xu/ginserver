package admin

import (
	"ginserver/internal/app/models"

	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Controller
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Log(models.LogTypeLogout, models.LogLevelInfo)
	p.RespRedirect302("/admin/login")
}
