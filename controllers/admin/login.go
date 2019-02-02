package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct{}

func (p *login) registerRouter(r *gin.RouterGroup) {
	r.GET("login", p.get)
	r.POST("login", p.post)
}

func (p *login) get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func (p *login) post(ctx *gin.Context) {
}
