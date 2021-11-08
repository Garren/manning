package main

import (
	"database/sql"
	"fmt"

	// this package has its own init() and will register the driver itself
	// on import. Otherwise we'd have to call
	//  sql.Register("postgres", &drv{})
	// the _ indicates that we're not to use the driver explicitly, instead we
	// rely on the database/sql interfaces it implements
	_ "github.com/lib/pq"
	// go get github.com/lib/pq
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=mypass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) { // named return
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return // named return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.
		QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) { //named return
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return // named return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return // named return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return //named return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return //named return
}

func main() {
	post := Post{Content: "Hello, World!", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
}
