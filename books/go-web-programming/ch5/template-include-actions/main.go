package main

import (
	"html/template"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("t1.html", "t2.html"))
	t.Execute(w, "hello world!")
}
func processMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmap.html"))
	m := map[string]int{"one": 1, "two": 2}
	t.Execute(w, m)
}
func processPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tpipe.html"))
	t.Execute(w, 12.3456)
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func processFunc(w http.ResponseWriter, r *http.Request) {
	// create a function map and register it with a template instance
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tfunc.html").Funcs(funcMap)
	// "fdate" is a function avaiable to our template.
	t = template.Must(t.ParseFiles("tfunc.html"))
	t.Execute(w, time.Now())
}

func processForm(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("form.html"))
	t.Execute(w, nil)
}

func processComment(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("comment.html"))
	t.Execute(w, r.FormValue("comment"))
}

func processLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("layout.html"))
	t.ExecuteTemplate(w, "layout", "")
}
func processRedLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("red-layout.html", "red-hello.html"))
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process-map", processMap)
	http.HandleFunc("/process-pipeline", processPipeline)
	http.HandleFunc("/process-func", processFunc)
	http.HandleFunc("/process-form", processForm)
	http.HandleFunc("/process-comment", processComment)
	http.HandleFunc("/process-layout", processLayout)
	http.HandleFunc("/process-red-layout", processRedLayout)
	server.ListenAndServe()
}
