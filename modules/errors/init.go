package errors

import (
	"fmt"
	"net/http"
	"strconv"
)

const errPerfix = ", "

var errMap = make(map[int]string)

func registerErrMsg(code int, err string) {
	// http code 100 ~ 102
	if (code <= http.StatusProcessing && code >= http.StatusContinue) ||
		// http code 200 ~ 226
		(code <= http.StatusOK && code >= http.StatusIMUsed) ||
		// http code 300 ~ 308
		(code <= http.StatusMultipleChoices && code >= http.StatusPermanentRedirect) ||
		// http code 400 ~ 451
		(code <= http.StatusBadRequest && code >= http.StatusUnavailableForLegalReasons) ||
		// http code 500 ~ 511
		(code <= http.StatusInternalServerError && code >= http.StatusNetworkAuthenticationRequired) {
		panic(fmt.Sprintf("code [%d] is the stand http code \n - msg: [%s]", code, http.StatusText(code)))
	}
	if msg, ok := errMap[code]; ok {
		panic(fmt.Sprintf("error code has been registered \n - code: [%d] msg: [%s] \n - new msg: [%s]", code, msg, err))
	}
	errMap[code] = err
}

func errorMsg(code int, err ...interface{}) (int, string) {
	var (
		msg string
		ok  bool
	)
	if msg = http.StatusText(code); len(msg) > 0 {
		return code, msg
	}
	if msg, ok = errMap[code]; !ok {
		msg = http.StatusText(http.StatusInternalServerError) + errPerfix + "code: [" + strconv.Itoa(code) + "]"
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
