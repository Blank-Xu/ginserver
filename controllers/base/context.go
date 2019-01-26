package base

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
	*Page
	Uid  int
	Role string
}

func (p *Context) CheckLogin() bool {
	return (p.Uid > 0) && (len(p.Role) > 0)
}
