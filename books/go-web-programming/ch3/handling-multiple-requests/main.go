package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// a handler is anything that implements the ServeHTTP interface
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

// a handler is anything that implements the ServeHTTP interface
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// use the default DefaultServeMux
	}
	// use the http module's handle function to attach handlers to our muxer
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
