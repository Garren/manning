package main

import (
	"fmt"
	"net/http"
)

// a handler function is anything that implements the interface below
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// a handler function is anything that implements the interface below
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// use the default DefaultServeMux
	}
	// use the http module's handlefunc routine to attach our handler functions.
	// it saves you the need to define a struct that implements ServerHTTP
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
