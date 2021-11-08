package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// combines the url and form values
	fmt.Fprintln(w, "Form")
	fmt.Fprintln(w, r.Form)

	// only provides the form values
	fmt.Fprintln(w, "PostForm")
	fmt.Fprintln(w, r.PostForm)
}

func processMultipart(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	// combines the url and form values
	fmt.Fprintln(w, "MultipartForm")
	fmt.Fprintln(w, r.MultipartForm)
}

func noProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.Form)

	fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process-multipart", processMultipart)
	http.HandleFunc("/no-process", noProcess)

	server.ListenAndServe()
}
