package admin

import (
	"ginserver/modules/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type logout struct{}

func (p *logout) registerRouter(r *gin.RouterGroup) {
	r.POST("logout", p.post)
}

func (p *logout) post(ctx *gin.Context) {
	if session := sessions.Default(ctx); session != nil {
		session.Clear()
	}
	e.RespRedirect308(ctx, redirectLocation)
}
