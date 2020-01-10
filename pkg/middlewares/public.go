package middlewares

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyUserId = `uid`
	KeyRoleId = `rid`
)

func GetUserId(ctx *gin.Context) int {
	return ctx.GetInt(KeyUserId)
}

// func SetUserId(ctx *gin.Context, userId int) {
// 	ctx.Set(KeyUserId, userId)
// }

func GetRoleId(ctx *gin.Context) int {
	return ctx.GetInt(KeyRoleId)
}

// func SetRoleId(ctx *gin.Context, roleId int) {
// 	ctx.Set(KeyUserId, roleId)
// }
