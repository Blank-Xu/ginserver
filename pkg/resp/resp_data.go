package resp

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

func RespRedirect301(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusMovedPermanently, location)
}

func RespRedirect302(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusFound, location)
}
