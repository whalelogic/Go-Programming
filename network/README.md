# Networking in Go

This guide provides an overview of common networking operations in Go, focusing on making HTTP requests and handling JSON data. The examples in the `fetch` and `json` subdirectories illustrate these concepts.

## 1. The `net/http` Package

The `net/http` package is the cornerstone of Go's networking capabilities for web applications. It provides a rich set of tools for building both HTTP clients and servers.

### Making HTTP GET Requests

The simplest way to make an HTTP GET request is to use `http.Get()`:

```go
import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    resp, err := http.Get("https://api.github.com/users/golang")
    if err != nil {
        // Handle error
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // Handle error
    }

    fmt.Println(string(body))
}
```

The `fetch/fetch.go` file demonstrates this basic pattern.

**Key Steps:**

1.  **`http.Get(url)`:** Sends an HTTP GET request to the specified URL. It returns an `*http.Response` and an `error`.
2.  **Error Handling:** Always check for errors after making a request.
3.  **`defer resp.Body.Close()`:** It is crucial to close the response body to prevent resource leaks. `defer` is the idiomatic way to ensure this happens.
4.  **`ioutil.ReadAll(resp.Body)`:** Reads the entire response body into a byte slice. For large responses, it's more memory-efficient to read the body in chunks.

### Customizing HTTP Requests

For more control over the request (e.g., setting headers, using other HTTP methods like POST), you can create a custom `http.Request` and use an `http.Client`:

```go
client := &http.Client{}
req, err := http.NewRequest("GET", "https://api.github.com", nil)
if err != nil {
    // Handle error
}

req.Header.Add("Accept", "application/vnd.github.v3+json")

resp, err := client.Do(req)
// ...
```

## 2. Working with JSON

JSON (JavaScript Object Notation) is the de facto standard for data exchange on the web. Go's standard library provides excellent support for encoding and decoding JSON through the `encoding/json` package.

### Decoding JSON (Unmarshaling)

"Unmarshaling" is the process of converting a JSON byte slice into a Go data structure (like a struct or a map).

The `json/fetch.go` and `json/main.go` files show how to fetch JSON data from an API and unmarshal it into a Go struct.

First, define a Go struct that matches the structure of the JSON data you expect. You can use struct tags to map JSON keys to struct fields.

```go
type User struct {
    Login     string `json:"login"`
    Name      string `json:"name"`
    PublicRepos int    `json:"public_repos"`
}
```

Then, use `json.Unmarshal()` to decode the JSON:

```go
var user User
err := json.Unmarshal(body, &user)
if err != nil {
    // Handle error
}

fmt.Printf("User: %+v\n", user)
```

Alternatively, you can use `json.NewDecoder()` to decode a JSON stream directly from an `io.Reader` (like an `http.Response.Body`), which is more memory-efficient:

```go
var user User
err := json.NewDecoder(resp.Body).Decode(&user)
```

### Encoding JSON (Marshaling)

"Marshaling" is the process of converting a Go data structure into a JSON byte slice.

```go
user := User{
    Login:     "gopher",
    Name:      "Go Gopher",
    PublicRepos: 99,
}

jsonData, err := json.Marshal(user)
if err != nil {
    // Handle error
}

fmt.Println(string(jsonData))
```

For pretty-printing the JSON, use `json.MarshalIndent()`:

```go
jsonData, err := json.MarshalIndent(user, "", "  ")
```

## 3. Error Handling in Networking

Network operations are inherently unreliable. It's essential to handle potential errors gracefully. Common errors include:

-   DNS resolution failures.
-   Connection timeouts.
-   Server errors (e.g., 500 Internal Server Error).
-   Client errors (e.g., 404 Not Found).

Always check the `error` value returned by `http.Get`, `client.Do`, etc. Also, check the HTTP status code of the response to ensure the request was successful.

```go
if resp.StatusCode != http.StatusOK {
    // Handle non-200 status code
}
```

## Conclusion

Go's standard library provides a powerful and ergonomic set of tools for networking. The `net/http` package makes it easy to build robust HTTP clients, while the `encoding/json` package simplifies working with JSON data. By combining these packages, you can easily interact with web services and build networked applications in Go.

