package base

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	userId int
	role   string
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

func (p *Context) Parse() error {
	session := sessions.Default(p.Context)
	if session != nil {
		vUid := session.Get("userId")
		vRole := session.Get("role")
		if vUid != nil && vRole != nil {
			var ok bool
			if p.userId, ok = vUid.(int); ok {
				if p.role, ok = vRole.(string); ok {
					return nil
				}
			}
		}
	}
	return errors.New("session is nil")
}

func (p *Context) GetRole() string {
	return p.role
}

func (p *Context) GetUserId() int {
	return p.userId
}

func (p *Context) IsLogin() bool {
	return (p.userId > 0) && (len(p.role) > 0)
}
