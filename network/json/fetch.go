package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchJSON(path string, dst interface{}) error {
	r, err := http.Get(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	buf, err := io.ReadAll(r.Request.Response.Body)
	if err != nil {
		fmt.Printf("reading r.Request.Response.Body: %s", buf)
	}
	return json.NewDecoder(r.Body).Decode(dst)
}
