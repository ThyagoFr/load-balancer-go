package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

var (
	servers      = []string{}
	currentIndex = 0
	client       = http.Client{Transport: &transport}
)

func init() {
	http.DefaultClient = &http.Client{Transport: &transport}
}

func ProcessRequests() {
	for {
		select {
		case req := <-requestCh:
			log.Println("request")
			if len(servers) == 0 {
				response, _ := json.Marshal(map[string]string{
					"message": "no server registered",
				})
				req.response.WriteHeader(http.StatusInternalServerError)
				req.response.Write(response)
				req.doneCh <- struct{}{}
				continue
			}
			currentIndex++
			if currentIndex == len(servers) {
				currentIndex = 0
			}
			host := servers[currentIndex]
			go ProcessRequest(host, req)
		}
	}

}

func ProcessRequest(host string, req *WebRequest) {
	uri, _ := url.Parse(req.request.URL.String())
	uri.Scheme = "http"
	uri.Host = host
	requestToHost, _ := http.NewRequest(req.request.Method, uri.String(), req.request.Body)
	for k, v := range Headers(req.request.Header) {
		requestToHost.Header.Add(k, v)
	}
	response, err := client.Do(requestToHost)
	if err != nil {
		req.response.WriteHeader(http.StatusInternalServerError)
		req.doneCh <- struct{}{}
		return
	}
	for k, v := range Headers(response.Header) {
		req.response.Header().Add(k, v)
	}
	io.Copy(req.response, response.Body)
	req.doneCh <- struct{}{}
}
