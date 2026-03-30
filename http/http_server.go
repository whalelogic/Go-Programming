// Package http provides a simple HTTP server implementation.
package http

import (
	"log"
	"net/http"
	"errors"
)


type HTTPServer struct {
	Addr string
	Handler http.Handler
	ErrChan chan error
}

func NewHTTPServer(addr string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		Addr: addr,
		Handler: handler,
		ErrChan: make(chan error, 1),
	}
}

func (s *HTTPServer) Start() {
	go func() {
		log.Printf("Starting HTTP server on %s\n", s.Addr)
		if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
			s.ErrChan <- err
		}
	}()
}

func (s *HTTPServer) WaitForShutdown() error {
	err := <-s.ErrChan
	if err != nil {
		return errors.New("HTTP server error: " + err.Error())
	}
	return nil
}

func (s *HTTPServer) Stop() {
	s.WaitForShutdown()
	log.Println("Stopping HTTP server")
}


func ServeData(addr string, data []string) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(data[0]))
	})
	server := NewHTTPServer(addr, handler)
	server.Start()
	return server.WaitForShutdown()
}



