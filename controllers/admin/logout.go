package admin

import (
	"github.com/gin-gonic/gin"
)

type logout struct {
	loginLocation string
}

func (p *logout) registerRouter(r *gin.RouterGroup, location string) {
	p.loginLocation = location
	r.GET("logout", p.get)
}

func (p *logout) get(ctx *gin.Context) {
	newCtx := NewContext(ctx)
	newCtx.SessionDestroy()
	newCtx.RespRedirect302(p.loginLocation)
}
