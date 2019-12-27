package http

import (
	"crypto/tls"
	"net/http"
	"time"
)



var (
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second*30,
	}
)


