package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var errSessionNil = errors.New("session is nil")

func SessionCreate(ctx *gin.Context, userId, roleId int) error {
	session := sessions.Default(ctx)
	if session == nil {
		return errSessionNil
	}

	session.Set(KeyUserId, userId)
	session.Set(KeyRoleId, roleId)
	if err := session.Save(); err != nil {
		return err
	}

	return nil
}

func SessionDestroy(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}

	session.Clear()

	return session.Save()
}

func SessionAuth(redirectLocation string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if session := sessions.Default(ctx); session != nil {
			// login session check
			if userId, ok := session.Get(KeyUserId).(int); ok {
				if roleId, ok := session.Get(KeyRoleId).(int); ok {
					if userId > 0 && roleId > 0 {
						SetSessionUserId(ctx, userId)
						SetSessionRoleId(ctx, roleId)
						return
					}
				}
			}
		}

		ctx.Redirect(http.StatusFound, redirectLocation)
		ctx.Abort()
	}
}

func GetSessionUserId(ctx *gin.Context) int {
	return ctx.GetInt(KeyUserId)
}

func SetSessionUserId(ctx *gin.Context, userId int) {
	ctx.Set(KeyUserId, userId)
}

func GetSessionRoleId(ctx *gin.Context) int {
	return ctx.GetInt(KeyUserId)
}

func SetSessionRoleId(ctx *gin.Context, roleId int) {
	ctx.Set(KeyUserId, roleId)
}
