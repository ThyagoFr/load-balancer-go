package http

import "net/http"

func Headers(headers http.Header) map[string]string {
	headResp := make(map[string]string)
	for key, value := range headers {
		var values string
		for _, headerValue := range value {
			values += headerValue + " "
		}
		headResp[key] = values
	}
	return headResp
}
