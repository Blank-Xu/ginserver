package context

import (
	"net/http"
)

// code 5000-5999
const (
	DatabaseErrCode = 5011
	DatabaseErrMsg  = "Database Error"
)

func (p *Context) AbortResponseInternalServerError(err ...interface{}) {
	const msg = "Internal Server Error"

	if len(err) > 0 {
		p.AddErrors(msg, err...)
	}

	p.AbortWithStatusJSON(
		http.StatusInternalServerError,
		Response{
			Code:   http.StatusInternalServerError,
			Msg:    msg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) ResponseDatabaseErr(err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(DatabaseErrMsg, err...)
	}

	p.JSON(
		http.StatusNotImplemented,
		Response{
			Code:   DatabaseErrCode,
			Msg:    DatabaseErrMsg,
			Errors: p.Errors,
		},
	)
}

func (p *Context) AbortResponseDatabaseErr(err ...interface{}) {
	if len(err) > 0 {
		p.AddErrors(DatabaseErrMsg, err...)
	}

	p.AbortWithStatusJSON(
		http.StatusNotImplemented,
		Response{
			Code:   DatabaseErrCode,
			Msg:    DatabaseErrMsg,
			Errors: p.Errors,
		},
	)
}
