package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello world!</h1></body>
</html>`
	// if we don't specify a content type, then the lib will try to determine the
	// type based on the first 500 bytes of the response
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // WriteHeader writes status codes. Headers are modified
	fmt.Fprintln(w, "No such service, try next door")
	// using the Header method on the ResponseWriter
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Adam Garren",
		Threads: []string{"one", "two", "three"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
