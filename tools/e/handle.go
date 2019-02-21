package e

import (
	"fmt"
	"net/http"
	"strconv"
)

func ErrorMsg(code int, err ...interface{}) (int, string) {
	var (
		msg string
		ok  bool
	)
	if msg, ok = errMap[code]; !ok {
		msg = http.StatusText(http.StatusInternalServerError) + errPerfix +
			"code: [" + strconv.Itoa(code) + "]"
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
