package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/whalelogic/Go-Programming/utils"
)



func main() {
	
	// Using the Swap function from the utils package
	x, y := utils.Swap("hello", 42)
	fmt.Println(x, y) 

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		h := r.Proto
		fmt.Fprintf(w, "Hello Protocol %s! Time is: %s\n", h, time.Now())
		// Response body: Hello Protocol HTTP/1.1! Time is: 2025-11-18 09:51:44.370398375 -0500 EST m=+30.972438668
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			message = "No message provided"
		}
		headers := r.Header
		for name, values := range headers {
			for _, value := range values {
				fmt.Fprintf(w, "Header: %s = %s\n", name, value)
			}
			// Response body example:
			// Header: User-Agent = curl/7.68.0
			// Header: Accept = */*
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Echo: %s", message)
		// Response body example:
		// If you visit: http://localhost:8080/echo?message=Hello,%20World!
		// Echo: Hello, World!
	})

	// Start the HTTP server
	fmt.Println("\n\n\t 🖧Server listening localhost:8080\n\n --> GET or Ping /greet or /echo \n\n --> Using HTTP/1.1 Protocol")
	http.ListenAndServe(":8080", nil)


}
