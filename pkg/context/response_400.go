package context

import (
	"net/http"
)

// code 4000-4999
const (
	InvalidParamsCode = 4000
	InvalidParamsMsg  = "Invalid Params"
)

func (p *Context) ResponseInvalidParams(err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(InvalidParamsMsg, err...)
	}

	p.JSON(
		http.StatusBadRequest,
		Response{
			Code:   InvalidParamsCode,
			Msg:    InvalidParamsMsg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) AbortResponseInvalidParams(err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(InvalidParamsMsg, err...)
	}

	p.AbortWithStatusJSON(
		http.StatusBadRequest,
		Response{
			Code:   InvalidParamsCode,
			Msg:    InvalidParamsMsg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) ResponseUnauthorized(err ...interface{}) {
	const msg = "Unauthorized"

	if len(err) > 0 {
		p.AddErrors(msg, err...)
	}

	p.JSON(
		http.StatusUnauthorized,
		Response{
			Code:   http.StatusUnauthorized,
			Msg:    msg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) AbortResponseUnauthorized(err ...interface{}) {
	const msg = "Unauthorized"

	if len(err) > 0 {
		p.AddErrors(msg, err...)
	}

	p.AbortWithStatusJSON(
		http.StatusUnauthorized,
		Response{
			Code:   http.StatusUnauthorized,
			Msg:    msg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) ResponseForbidden() {
	p.JSON(
		http.StatusForbidden,
		Response{
			Code: http.StatusForbidden,
			Msg:  "Forbidden",
		},
	)
}

func (p *Context) AbortResponseForbidden() {
	p.AbortWithStatusJSON(
		http.StatusForbidden,
		Response{
			Code: http.StatusForbidden,
			Msg:  "Forbidden",
		},
	)
}

func (p *Context) ResponseNotFound() {
	p.JSON(
		http.StatusNotFound,
		Response{
			Code: http.StatusNotFound,
			Msg:  "Not Found",
		},
	)
}

func (p *Context) AbortResponseNotFound() {
	p.AbortWithStatusJSON(
		http.StatusNotFound,
		Response{
			Code: http.StatusNotFound,
			Msg:  "Not Found",
		},
	)
}

func (p *Context) ResponseMethodNotAllowed() {
	p.JSON(
		http.StatusMethodNotAllowed,
		Response{
			Code: http.StatusMethodNotAllowed,
			Msg:  "Method Not Allowed",
		},
	)
}

func (p *Context) AbortResponseMethodNotAllowed() {
	p.AbortWithStatusJSON(
		http.StatusMethodNotAllowed,
		Response{
			Code: http.StatusMethodNotAllowed,
			Msg:  "Method Not Allowed",
		},
	)
}
