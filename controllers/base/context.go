package base

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
	*Conditional
	Uid  int
	Role string
}

func (p *Context) IsLogin() bool {
	return (p.Uid > 0) && (len(p.Role) > 0)
}
