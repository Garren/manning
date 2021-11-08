package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"git.sr.ht/~garren/go-web-programming/chitchat/data"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}) {
	fmt.Println(a...)
}

// executed by the runtime after the package has been imported.
// https://golangdocs.com/init-function-in-golang
func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json") // open for read access only
	if err != nil {
		log.Fatalln("Cannot open config file", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
		os.Exit(1)
	}
}

// helper - redirect to error message page
func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), http.StatusFound)
}

// get a session for the uuid in the request's cookie or fail
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return // note: named return variables
}

// instanciate a template for an array of template files
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return // named return
}

// generate html for an array of template files, write to response
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO")
	logger.Println(args...)
}
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}
func warning(args ...interface{}) {
	logger.SetPrefix("WARNING")
	logger.Println(args...)
}

// version
func version() string {
	return "0.1"
}
