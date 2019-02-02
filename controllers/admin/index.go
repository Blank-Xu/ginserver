package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type index struct{}

func (p *index) registerRouter(r *gin.RouterGroup) {
	r.GET("/", p.Get)
}

func (p *index) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.html", nil)
}
