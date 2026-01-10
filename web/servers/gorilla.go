package servers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func NewServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla 🦍!")
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	return srv
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla 🦍!")
	})

	http.ListenAndServe(":8080", r)
}
