package e

import (
	"fmt"
	"net/http"
	"strconv"
)

func errorMsg(code errCode, err ...interface{}) (errCode, string) {
	var (
		msg string
		ok  bool
	)
	if msg, ok = errMap[code]; !ok {
		msg = http.StatusText(http.StatusInternalServerError) + errPerfix +
			"code: [" + strconv.Itoa(int(code)) + "]"
		code = http.StatusInternalServerError
	}
	if len(err) > 0 {
		switch str := err[0].(type) {
		case string:
			msg += errPerfix + str
		case error:
			msg += errPerfix + str.Error()
		default:
			msg += errPerfix + fmt.Sprintf("%v", str)
		}
	}
	return code, msg
}

type codeMsg struct {
	Code errCode `json:"code"`
	Msg  string  `json:"msg"`
}

func NewCodeMsg(code errCode, msg string) *codeMsg {
	return &codeMsg{Code: code, Msg: msg}
}
