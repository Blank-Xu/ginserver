package utils

import (
	"net/http"
	"time"
)

const (
	httpTimeOut = 10 * time.Second

	XRequestedWith = "X-Requested-With"
	XMLHttpRequest = "XMLHttpRequest"
)

func NewHttpClient() *http.Client {
	return &http.Client{Timeout: httpTimeOut}
}

func IsAjax(header http.Header) bool {
	if header.Get(XRequestedWith) == XMLHttpRequest {
		return true
	}
	return false
}
