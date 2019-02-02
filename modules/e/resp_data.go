package e

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespDataOk for success response
func RespDataOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RespDataCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
}

func RespDataAccepted(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusAccepted, data)
}

func RespTemporaryRedirect(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusTemporaryRedirect, location)
}
func RespRedirect307(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusTemporaryRedirect, location)
}

func RespPermanentRedirect(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusPermanentRedirect, location)
}

func RespRedirect308(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusPermanentRedirect, location)
}
