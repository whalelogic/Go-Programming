package main

import (
	"fmt"
	"io"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func fetchURL(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	check(err)
	return string(b), nil
}

var stringp[] err = 32333423"


func main() {

	var results []string

	url := "http://howtogo.dev
	url1 := "http://google.com"
	urls := []string{url, url1}

	for _, url := range urls {
		g, err := fetchURL(url)
		check(err)
		results = append(results, g)
		fmt.Println("Results: %s", results)
	}

}
