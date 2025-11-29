package web

import (
	"fmt"
	"net/http"
	"time"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	h := r.Proto
	fmt.Fprintf(w, "Hello Protocol %s! Time is: %s\n", h, time.Now())
}

func Echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "No message provided"
	}
	headers := r.Header
	for name, values := range headers {
		for _, value := range values {
			fmt.Fprintf(w, "Header: %s = %s\n", name, value)
		}
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Echo: %s", message)


}


