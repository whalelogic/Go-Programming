package main

import (
	"fmt"
	"net/http"
	"github.com/whalelogic/Go-Programming/web"
)


func main() {
	r := servers.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla 🦍!")
	})

	http.ListenAndServe(":8080", r)
}
