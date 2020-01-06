package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/pkg/resp"
)

type Controller struct {
	*gin.Context
}

func (p *Controller) New(context *gin.Context) {
	p.Context = context
}

// RespOk 200
// for GET
func (p *Controller) RespOk(data interface{}) {
	if data == nil {
		p.AbortWithStatus(http.StatusOK)
		return
	}
	p.AbortWithStatusJSON(http.StatusOK, data)
}

// RespCreated 201
// for POST/PUT/PATCH
func (p *Controller) RespCreated(data interface{}) {
	if data == nil {
		p.AbortWithStatus(http.StatusCreated)
		return
	}
	p.AbortWithStatusJSON(http.StatusCreated, data)
}

// RespAccepted  202
// for async task
func (p *Controller) RespAccepted() {
	p.AbortWithStatus(http.StatusAccepted)
}

// RespNoContent 204
// for DELETE
func (p *Controller) RespNoContent() {
	p.AbortWithStatus(http.StatusNoContent)
}

func (p *Controller) RespErrInvalidParams(err ...interface{}) {
	p.AbortWithStatusJSON(http.StatusBadRequest, resp.RespErrCode(resp.CodeInvalidParams, err...))
}

func (p *Controller) RespErrForbidden() {
	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusForbidden))
}

func (p *Controller) RespErrNotFound() {
	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusNotFound))
}

func (p *Controller) RespErrInternalServerError(err error) {
	p.Error(err)
	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusInternalServerError))
}

func (p *Controller) RespErrDBError(err error) {
	p.Error(err)
	if gin.Mode() != gin.ReleaseMode {
		p.AbortWithStatusJSON(http.StatusNotImplemented, resp.RespErrCode(resp.CodeDBErr, err))
	} else {
		p.AbortWithStatusJSON(http.StatusNotImplemented, resp.RespErrCode(resp.CodeDBErr))
	}
}

// RespErrCode return http code, server code and msg struct
func (p *Controller) RespErrCode(code int, err ...interface{}) *resp.ResponseErr {
	return resp.NewResponseErr(resp.ErrorMsg(code, err...))
}

//  RespErrHttp return http code and msg struct
func (p *Controller) RespErrHttp(httpCode int) (int, *resp.ResponseErr) {
	return httpCode, resp.NewResponseErr(httpCode, http.StatusText(httpCode))
}
