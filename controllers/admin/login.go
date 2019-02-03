package admin

import (
	"ginserver/modules/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct{}

func (p *login) registerRouter(r *gin.RouterGroup) {
	r.GET("login", p.get)
	r.POST("login", p.post)
}

func (p *login) get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"Title": "ginserverLogin",
		})
}

func (p *login) post(ctx *gin.Context) {
	newCtx := NewContextLogin(ctx, 1, "admin")
	if err := newCtx.CreateSession(); err != nil {
		e.RespErrInternalServerError(ctx)
		return
	}
	e.RespRedirect302(ctx, "/admin")
}
