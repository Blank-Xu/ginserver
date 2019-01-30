package e

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/modules/resp"
)

// RespErrCode return http code, server code and msg struct
func RespErrCode(code errCode, err ...interface{}) *resp.ResponseErr {
	code, msg := errorMsg(code, err...)
	return resp.NewResponseErr(NewCodeMsg(code, msg))
}

//  RespErrHttp return http code and msg struct
func RespErrHttp(httpCode int) (int, *resp.ResponseErr) {
	msg := http.StatusText(httpCode)
	return httpCode, resp.NewResponseErr(NewCodeMsg(errCode(httpCode), msg))
}

func RespErrParamsInvalid(c *gin.Context, err ...interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, RespErrCode(CodeParamInvalid, err...))
}

func RespErrNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(RespErrHttp(http.StatusNotFound))
}

func RespErrDBError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, RespErrCode(CodeDBErr, err))
}

func RespSuccOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resp.NewResponseData(data))
}

func RespSuccCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, resp.NewResponseData(data))
}

func RespSuccAccepted(c *gin.Context, data interface{}) {
	c.JSON(http.StatusAccepted, resp.NewResponseData(data))
}
