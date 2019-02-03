package admin

import (
	"ginserver/modules/e"

	"github.com/gin-contrib/sessions"
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
	if session := sessions.Default(ctx); session != nil {
		session.Clear()
	}
	e.RespRedirect302(ctx, p.loginLocation)
}
