package middlewares

import (
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetJwtUserId(ctx *gin.Context) int {
	return ctx.GetInt(KeyUserId)
}

func SetJwtUserId(ctx *gin.Context, userId int) {
	ctx.Set(KeyUserId, userId)
}

func GetJwtRoleId(ctx *gin.Context) int {
	return ctx.GetInt(KeyRoleId)
}

func SetJwtRoleId(ctx *gin.Context, roleId int) {
	ctx.Set(KeyRoleId, roleId)
}
