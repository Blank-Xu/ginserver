package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	uid  int
	role string
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

func NewContextLogin(context *gin.Context, uid int, role string) *Context {
	return &Context{Context: context, uid: uid, role: role}
}

func (p *Context) CreateSession() error {
	session := sessions.Default(p.Context)
	session.Set("uid", p.uid)
	session.Set("role", p.role)
	return session.Save()
}

func (p *Context) ParseSession() (ok bool) {
	session := sessions.Default(p.Context)
	if session != nil {
		vUid := session.Get("uid")
		vRole := session.Get("role")
		if vUid != nil && vRole != nil {
			if p.uid, ok = vUid.(int); ok {
				if p.role, ok = vRole.(string); ok {
					return
				}
			}
		}
	}
	return
}

func (p *Context) GetRole() string {
	return p.role
}

func (p *Context) GetUserId() int {
	return p.uid
}

func (p *Context) IsLogin() bool {
	return (p.uid > 0) && (len(p.role) > 0)
}
