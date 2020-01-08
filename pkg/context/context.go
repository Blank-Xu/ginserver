package context

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"ginserver/pkg/middlewares"
)

const (
	MaxPageSize = 100
)

func New(ctx *gin.Context) *Context {
	// todo:
	userId := middlewares.GetSessionUserId(ctx)
	roleId := middlewares.GetSessionRoleId(ctx)

	return &Context{
		Context: ctx,
		UserId:  userId,
		RoleId:  roleId}
}

type Context struct {
	*gin.Context
	UserId int
	RoleId int
}

func (p *Context) GetPage() (page, pageSize int) {
	page, _ = strconv.Atoi(
		p.DefaultQuery("page", "1"))

	pageSize, _ = strconv.Atoi(
		p.DefaultQuery("page_size", "5"))
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	return
}

func (p *Context) AddErrors(meta interface{}, err ...interface{}) {
	l := len(err)
	if l == 0 {
		return
	}

	errs := make([]*gin.Error, l)
	for k, v := range err {
		switch v.(type) {
		case error:
			errs[k] = &gin.Error{
				Err:  err[k].(error),
				Type: gin.ErrorTypeAny,
				Meta: meta,
			}
		case string:
			errs[k] = &gin.Error{
				Err:  errors.New(err[k].(string)),
				Type: gin.ErrorTypeAny,
				Meta: meta,
			}
		}
	}

	p.Errors = append(p.Errors, errs...)
}

func (p *Context) Response(httpCode, code int, msg string, err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(msg, err)
	}

	p.JSON(
		httpCode,
		Response{
			Code:   code,
			Msg:    msg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) AbortResponse(httpCode, code int, msg string, err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(msg, err)
	}

	p.AbortWithStatusJSON(
		httpCode,
		Response{
			Code:   code,
			Msg:    msg,
			Errors: p.Errors,
		},
	)
}
