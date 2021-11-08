package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=mypass sslmode=disable")
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/*
$ http :8080/post/4
HTTP/1.1 200 OK
Content-Length: 70
Content-Type: application/json
Date: Sun, 08 Aug 2021 04:24:01 GMT

{
    "author": "Sau Sheong",
    "content": "My first post!",
    "id": 4
}
*/
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = post.fetch(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(post, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// $ http -v :8080/post/ content="My first post\!" author="Sau Sheong"
func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	json.Unmarshal(body, post)
	err = post.create()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

/*
$ http PUT :8080/post/4 content="Updated Content" author="Sau Sheong"
HTTP/1.1 200 OK
Content-Length: 0
Date: Sun, 08 Aug 2021 04:26:37 GMT
*/
func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = post.fetch(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	json.Unmarshal(body, post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

/*
$ http DELETE :8080/post/3
HTTP/1.1 200 OK
Content-Length: 0
Date: Sun, 08 Aug 2021 04:27:40 GMT
*/
func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = post.fetch(id)
	if err != nil {
		return
	}

	err = post.delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}
