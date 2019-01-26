package errors

import (
	"net/http"

	"ginserver/modules/resp"
)

type codeMsg struct {
	Code errCode `json:"code"`
	Msg  string  `json:"msg"`
}

// RespError return http code, server code and msg struct
func RespError(httpCode int, code errCode, err ...interface{}) (int, interface{}) {
	code, msg := errorMsg(code, err...)
	return httpCode,
		&resp.ResponseErr{
			Error: &codeMsg{
				Code: code,
				Msg:  msg,
			}}
}

//  RespHttpError return http code and msg struct
func RespHttpError(httpCode int) (int, interface{}) {
	return httpCode,
		&resp.ResponseErr{
			Error: &codeMsg{
				Code: errCode(httpCode),
				Msg:  http.StatusText(httpCode),
			}}
}
