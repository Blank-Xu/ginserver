package e

import (
	"net/http"

	"ginserver/modules/resp"
)

type codeMsg struct {
	Code errCode `json:"code"`
	Msg  string  `json:"msg"`
}

// RespError return http code, server code and msg struct
func RespError(code errCode, err ...interface{}) *resp.ResponseErr {
	code, msg := errorMsg(code, err...)
	return &resp.ResponseErr{
		Error: &codeMsg{
			Code: code,
			Msg:  msg,
		}}
}

//  RespHttpError return http code and msg struct
func RespHttpError(httpCode int) (int, *resp.ResponseErr) {
	return httpCode,
		&resp.ResponseErr{
			Error: &codeMsg{
				Code: errCode(httpCode),
				Msg:  http.StatusText(httpCode),
			}}
}
