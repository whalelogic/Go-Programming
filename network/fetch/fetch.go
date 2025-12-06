// Package fetch fetches the content of a url and prints it to the console
package fetch 

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"os"
)


func main() {

	// check if there are command line arguments
	if len(os.Args) < 2 {
		log.Print("Must enter at least one url")
		os.Exit(1)
	}

	// iterate over command line arguments
	for _, url := range os.Args[1:] {
		if url == "" {
			log.Print("Must enter a url")
		}
		r, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching url: %v", err)
			os.Exit(1)
		}

		// create a buffer to read the response body into
		buf, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %s: %v\n", url, err)
			os.Exit(1)
		}
		// fetch prints to the console the http response body
		fmt.Printf("%s", buf)
	}
	
}
