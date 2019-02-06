package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (p *Context) SessionCreate() error {
	session := sessions.Default(p.Context)
	if session != nil {
		session.Set("uid", p.uid)
		session.Set("role", p.role)
		return session.Save()
	}
	return errors.New("session is nil")
}

func (p *Context) SessionParse() (ok bool) {
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

func (p *Context) SessionDestroy() {
	session := sessions.Default(p.Context)
	if session != nil {
		session.Clear()
		if err := session.Save(); err != nil {
			p.Error(err)
		}
	}
}

func (p *Context) GetRole() string {
	return p.role
}

func (p *Context) GetUserId() int {
	return p.uid
}
