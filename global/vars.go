package global

import (
	"crypto/tls"
	"net/http"
)

var (
	AppName string
	RunMode string
)

var (
	DefaultHttpClient = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: DefaultHttpTimeOutSecond,
	}
)
