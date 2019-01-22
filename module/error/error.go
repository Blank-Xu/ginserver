package error

import (
	"fmt"
	"strconv"
)

const errPerfix = ", "

var errMap = make(map[int]string)

func regErrCode(code int, err string) {
	if msg, ok := errMap[code]; ok {
		panic(fmt.Sprintf("code: [%d], msg: [%s], has registered, error: [%s]", code, msg, err))
	}
	errMap[code] = err
}

func Error(code int, err ...interface{}) (int, string) {
	var (
		msg string
		ok  bool
	)
	if msg, ok = errMap[code]; !ok {
		msg = Code500Err + errPerfix + "code: " + strconv.Itoa(code)
		code = Code500
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
