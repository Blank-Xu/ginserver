package e

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseErr for failed response
type ResponseErr struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"status bad request"`
}

func NewResponseErr(code int, msg string) *ResponseErr {
	return &ResponseErr{
		Code: code,
		Msg:  msg,
	}
}

// RespErrCode return http code, server code and msg struct
func RespErrCode(code int, err ...interface{}) *ResponseErr {
	return NewResponseErr(ErrorMsg(code, err...))
}

//  RespErrHttp return http code and msg struct
func RespErrHttp(httpCode int) (int, *ResponseErr) {
	return httpCode, NewResponseErr(httpCode, http.StatusText(httpCode))
}

func RespErrInvalidParams(ctx *gin.Context, err ...interface{}) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, RespErrCode(CodeInvalidParams, err...))
}

func RespErrForbidden(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(RespErrHttp(http.StatusForbidden))
}

func RespErrNotFound(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(RespErrHttp(http.StatusNotFound))
}

func RespErrInternalServerError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(RespErrHttp(http.StatusInternalServerError))
}

func RespErrDBError(ctx *gin.Context, err error) {
	ctx.Error(err)
	if gin.Mode() != gin.ReleaseMode {
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, RespErrCode(CodeDBErr, err))
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, RespErrCode(CodeDBErr))
	}
}
