package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var errSessionNil = errors.New("session is nil")

func SessionCreate(context *gin.Context, userId, roleId int) error {
	if session := sessions.Default(context); session != nil {
		session.Set(KeyUserId, userId)
		session.Set(KeyRoleId, roleId)
		if err := session.Save(); err != nil {
			return err
		}
		return nil
	}
	return errSessionNil
}

func SessionDestroy() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session := sessions.Default(c); session != nil {
			session.Clear()
			if err := session.Save(); err != nil {
				c.Error(err)
			}
		}
	}
}

func SessionAuth(redirectLocation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if session := sessions.Default(c); session != nil {
			// login session check
			uid := session.Get(KeyUserId)
			role := session.Get(KeyRoleId)
			if userId, ok := uid.(int); ok {
				if roleId, ok := role.(int); ok {
					if userId > 0 && roleId > 0 {
						c.Set(KeyUserId, userId)
						c.Set(KeyRoleId, roleId)
						return
					}
				}
			}
		}
		c.Redirect(http.StatusFound, redirectLocation)
		c.Abort()
	}
}
