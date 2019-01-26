package errors

import (
	"net/http"

	"ginserver/modules/resp"
)

// ErrorCodeJson return http code, server code and msg
// include http status code
func ErrorCodeJson(httpCode, code int, err ...interface{}) (int, *resp.Response) {
	code, msg := errorMsg(code, err...)
	return httpCode, &resp.Response{Code: code, Msg: msg}
}

//  ErrorHttpCodeJson return http code and msg
func ErrorHttpCodeJson(httpCode int) (int, *resp.Response) {
	return httpCode, &resp.Response{Code: httpCode, Msg: http.StatusText(httpCode)}
}
