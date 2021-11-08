package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)

	encoding := r.Header.Get("Accept-Encoding")
	fmt.Fprintf(w, "Got %s\n", encoding)

	lang := r.Header["Accept-Language"]
	fmt.Fprintf(w, "Got %s\n", lang)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
