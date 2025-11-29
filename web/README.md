# Building Web Applications in Go

This guide provides an overview of how to build web servers and clients in Go using the powerful `net/http` package from the standard library. The `server.go` and `client.go` files in this directory serve as practical examples of these concepts.

## 1. The `net/http` Package

The `net/http` package contains all the tools you need to work with the HTTP protocol in Go. It's comprehensive, robust, and designed to be easy to use for building production-ready web services.

## 2. Creating a Web Server (`server.go`)

Building a web server in Go is remarkably simple. The `server.go` file demonstrates the core components of a basic Go web server.

### Key Components:

1.  **Handler Functions:** A handler is a function that receives an `http.ResponseWriter` and an `http.Request`. It's responsible for processing the request and writing a response.

    A handler function has the signature: `func(w http.ResponseWriter, r *http.Request)`

    -   `http.ResponseWriter`: An interface used to construct the HTTP response. You write your headers, status code, and response body to it.
    -   `http.Request`: A struct that contains all the information about the incoming HTTP request, including the URL, headers, and body.

    ```go
    func helloHandler(w http.ResponseWriter, r *http.Request) {
        // Set the Content-Type header
        w.Header().Set("Content-Type", "text/plain")

        // Write the response body
        fmt.Fprintln(w, "Hello, World!")
    }
    ```

2.  **Routing:** The `http.HandleFunc()` function is used to register a handler function for a given URL path pattern. This is how you tell the server which function to call for which URL.

    ```go
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/about", aboutHandler)
    ```

3.  **Starting the Server:** The `http.ListenAndServe()` function starts an HTTP server on a given address and port. It blocks until the server is shut down.

    ```go
    fmt.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    ```
    The second argument to `ListenAndServe` is a handler. If it's `nil`, the default serve mux (`http.DefaultServeMux`) is used, which is where `http.HandleFunc` registers its handlers.

### Running the Server

To run the server, you would execute:

```sh
go run server.go
```

You can then access the server in your browser or with a tool like `curl`:

```sh
curl http://localhost:8080
# Output: Hello, World!
```

## 3. Creating a Web Client (`client.go`)

The `net/http` package also makes it trivial to act as an HTTP client, sending requests to other servers. The `client.go` file shows a basic implementation.

### Making a GET Request

The simplest way to make a GET request is with `http.Get()`.

```go
resp, err := http.Get("http://localhost:8080")
if err != nil {
    // Handle error
    log.Fatalf("Failed to get URL: %v", err)
}
defer resp.Body.Close()

// Check the status code
if resp.StatusCode != http.StatusOK {
    log.Fatalf("Request failed with status: %s", resp.Status)
}

// Read the response body
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    log.Fatalf("Failed to read response body: %v", err)
}

fmt.Println("Response:", string(body))
```

**Important:** Always remember to `defer resp.Body.Close()` to prevent resource leaks.

### Customizing Requests

For more control (e.g., making POST requests, setting headers), you can create a custom `http.Client` and `http.Request`.

```go
client := &http.Client{
    Timeout: 10 * time.Second, // Set a timeout
}

req, err := http.NewRequest("POST", "http://example.com/upload", someReader)
if err != nil {
    // Handle error
}
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
// ...
```

The `fetch_url` executable in this directory is likely a compiled version of a client program that fetches content from a URL.

## 4. Advanced Server Concepts

-   **`http.ServeMux`:** For more complex applications, instead of using the default serve mux, you can create your own instance of `http.ServeMux`. This allows for better separation of concerns and easier testing.

    ```go
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", mux)
    ```

-   **Middleware:** Middleware (or chaining handlers) is a powerful pattern for encapsulating common functionality like logging, authentication, or compression. A middleware is essentially a handler that wraps another handler.

-   **Serving Static Files:** The `http.FileServer` handler can be used to serve static files (like CSS, JavaScript, and images) from a directory.

    ```go
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    ```

## Conclusion

Go's standard library provides a first-class experience for web development. With the `net/http` package, you can build high-performance, production-grade web servers and clients with minimal boilerplate. The simplicity and power of the standard library are key reasons why Go is a popular choice for building APIs and web services.
