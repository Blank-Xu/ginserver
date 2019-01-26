package util

import (
	"net/http"
	"time"
)

const (
	httpTimeOut = 15 * time.Second
)

func NewHttpClient() *http.Client {
	return &http.Client{Timeout: httpTimeOut}
}
