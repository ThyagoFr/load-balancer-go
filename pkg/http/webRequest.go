package http

import "net/http"

type WebRequest struct {
	request  *http.Request
	response http.ResponseWriter
	doneCh   chan struct{}
}
