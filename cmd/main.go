package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		requestCh <- &WebRequest{request: r, response: w, doneCh: doneCh}
		<-doneCh
	})

	go processRequests()

	go http.ListenAndServe(":2000", nil)
	log.Printf("Server started, press <ENTER> to exit")
	fmt.Scan()

}
