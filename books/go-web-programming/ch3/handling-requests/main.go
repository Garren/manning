package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// a handler is anything that implements the ServeHTTP interface
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, // we assign a single handler, no longer using muxers
	}
	server.ListenAndServe()
}
