// Package web contains web relatd utilities.
package clients

import (
	"io"
	"fmt"
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

func PostURL(url string, data string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	baseURL :=  "https://api.example.com/search"

	get := http.MethodGet
	req, _ := http.NewRequest(get, baseURL, nil)
	q := req.URL.Query()

	q.Add("search", "keiththomson.dev")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	var b io.ReadAll
	res, _ := http.NewRequest(get, baseURL, b)




}



