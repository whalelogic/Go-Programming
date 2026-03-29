package utilities 

import (
	"fmt"
	"log"
	"net/url"
)

// ParseURL takes a URL string and returns its components.

func ParseURL(URL string) (*url.URL, error) {
	parsed, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	fmt.Println("Parsed URL:", parsed)
	fmt.Println("Scheme: ", parsed.Scheme)
	log.Println("Host: ", parsed.Host)
	fmt.Println("Path: ", parsed.Path)
	return parsed, nil
}
