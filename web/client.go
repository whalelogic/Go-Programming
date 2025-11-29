// Package web creates functions, methods and demonstrates those defined in Go's net/http package 
package web

import (
	"net/http"
	"time"
)

func GetURL(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

