// Package json contains functions to fetch and decode JSON data from a given URL.
package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchJSON(path string, dst any) error {
	r, err := http.Get(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	// Ensure the response body is closed after reading
	defer r.Body.Close()

	// Read the response body into a buffer
	buf, err := io.ReadAll(r.Request.Response.Body)
	if err != nil {
		fmt.Printf("Error reading r.Request.Response.Body: %s", buf)
	}
	return json.NewDecoder(r.Body).Decode(dst)
}

// GetJSON is same func as above, just with different return type
// This time the func returns map[string]any instead of reading into a buffer before decoding
func GetJSON(path string) (map[string]any, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
