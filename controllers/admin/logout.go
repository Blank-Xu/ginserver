package admin

import (
	"github.com/rs/zerolog"

	"ginserver/models/log"

	"github.com/gin-gonic/gin"
)

type ControllerLogout struct {
	Controller
}

func (p *ControllerLogout) Get(ctx *gin.Context) {
	p.New(ctx)
	p.LogDB(log.TypeLogout, zerolog.InfoLevel)
	p.RespRedirect302("/admin/login")
}
