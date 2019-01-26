package error

import (
	"fmt"
	"net/http"
	"strconv"
)

const errPerfix = ", "

var errMap = make(map[int]string)

func registerErrMsg(code int, err string) {
	if msg, ok := errMap[code]; ok {
		panic(fmt.Sprintf("error code has been registered \n - code: [%d] msg: [%s] \n - new msg: [%s]", code, msg, err))
	}
	errMap[code] = err
}

func ErrorMsg(code int, err ...interface{}) (int, string) {
	var (
		msg string
		ok  bool
	)
	if msg, ok = errMap[code]; !ok {
		msg = http.StatusText(http.StatusInternalServerError) + errPerfix + "code: " + strconv.Itoa(code)
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

func ErrorMap(code int, err ...interface{}) map[int]string {
	code, msg := ErrorMsg(code, err...)
	return map[int]string{code: msg}
}
