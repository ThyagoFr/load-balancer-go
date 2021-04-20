package http

import (
	"crypto/tls"
	"net/http"
)

var (
	transport = http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)
