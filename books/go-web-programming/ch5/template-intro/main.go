package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello world!")
}

func random(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmpl-rand.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmpl-iter.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func action(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmpl-action.html"))
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/random", random)
	http.HandleFunc("/iterate", iterate)
	http.HandleFunc("/action", action)
	server.ListenAndServe()
}
