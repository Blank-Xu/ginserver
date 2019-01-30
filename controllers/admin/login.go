package admin

import (
	"github.com/gin-gonic/gin"
)

type Login struct{}

func (p *Login) RegisterRouter(r *gin.RouterGroup) {
	r.GET("login", p.Get)
	r.POST("login", p.Post)
}

func (p *Login) Get(c *gin.Context) {

}

func (p *Login) Post(c *gin.Context) {

}
