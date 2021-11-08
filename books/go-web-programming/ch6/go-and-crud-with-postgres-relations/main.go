package main

import (
	"database/sql"
	"errors"
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
	Id       int
	Content  string
	Author   string
	Comments []Comment // note: slices are actually pointers
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=mypass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return // named return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.
		QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return
	}

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
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

func main() {
	post := Post{Content: "Hello, World!", Author: "Sau Sheong"}
	post.Create()

	comment := Comment{Content: "Good point!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, err := GetPost(post.Id)
	if err != nil {
		return
	}

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
