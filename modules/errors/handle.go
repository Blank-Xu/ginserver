package errors

import (
	"net/http"
)

type respErrorCodeJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ErrorCodeJson return code and respErrorCodeJson pointer struct
// include http status code
func ErrorCodeJson(httpCode, code int, err ...interface{}) (int, *respErrorCodeJson) {
	code, msg := errorMsg(code, err...)
	return httpCode, &respErrorCodeJson{code, msg}
}

//  ErrorHttpCodeJson return http code and respErrorCodeJson pointer struct
func ErrorHttpCodeJson(httpCode int) (int, *respErrorCodeJson) {
	return httpCode, &respErrorCodeJson{httpCode, http.StatusText(httpCode)}
}
