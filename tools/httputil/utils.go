package httputil

import (
	"net/http"
)

const (
	XRequestedWith = "X-Requested-With"
	XMLHttpRequest = "XMLHttpRequest"
)

func IsAjax(header http.Header) bool {
	if header.Get(XRequestedWith) == XMLHttpRequest {
		return true
	}
	return false
}
