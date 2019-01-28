package e

import (
	"fmt"
	"net/http"
)

type errCode int

const errPerfix = ", "

var errMap = make(map[errCode]string)

func registerErrMsg(code errCode, err string) {
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
		panic(fmt.Sprintf("code [%d] is the stand http code \n - msg: [%s]", code, http.StatusText(int(code))))
	}
	if msg, ok := errMap[code]; ok {
		panic(fmt.Sprintf("error code has been registered \n - code: [%d] msg: [%s] \n - new msg: [%s]", code, msg, err))
	}
	errMap[code] = err
}
