package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// a handler function is anything that implements the interface below
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func protect(h http.HandlerFunc) http.HandlerFunc {
	// return a handler function that calls our handlerfunc argument in a
	// closure. Log the name of the wrapped handlerfunc via reflection
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protecting....")
		h(w, r)
	}
}

func log(h http.HandlerFunc) http.HandlerFunc {
	// return a handler function that calls our handlerfunc argument in a
	// closure. Log the name of the wrapped handlerfunc via reflection
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// use the default DefaultServeMux
	}
	// use the http module's handlefunc routine to attach our handler functions.
	// it saves you the need to define a struct that implements ServerHTTP
	http.HandleFunc("/hello", protect(log(hello)))

	server.ListenAndServe()
}
