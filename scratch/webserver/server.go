package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)


func customHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Name", "Keith Thomson")
		next.ServeHTTP(w, r)
	})
}

func main() {

	str := "Hello from Go!"

	var b = []byte{75}
	fmt.Println("b: ", b)
	var ru rune
	fmt.Println("rune empty: ", ru)
	ru = '0'
	fmt.Println("rune A = 0: ", ru)
	ru = 'b'
	fmt.Println("rune B = b: ", ru)

	runeMap := make(map[rune]int)
	for _, v := range str {
		runeMap[v]++
	}
	fmt.Println("runeMap: ", runeMap)

	result := strings.Map(func(r rune) rune {
        return r + 1 // Shifts 'h' to 'i', 'e' to 'f', etc.
    }, str)
	fmt.Println("result: ", result)
	
	r := mux.NewRouter()
	
	// Apply the middleware to all routes
	r.Use(customHeaderMiddleware)
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var by = []byte(result)
		w.Write(by)
	})
	
	fmt.Printf("Server running on: %s", "localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))


}
