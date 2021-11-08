package main

import (

	// this package has its own init() and will register the driver itself
	// on import. Otherwise we'd have to call
	//  sql.Register("postgres", &drv{})
	// the _ indicates that we're not to use the driver explicitly, instead we
	// rely on the database/sql interfaces it implements
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// go get github.com/lib/pq
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=mypass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.
		QueryRowx("select id, content, author from posts where id = $1", id).
		StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) { //named return
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return // named return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.AuthorName).Scan(&post.Id)
	return // named return
}

func main() {
	post := Post{Content: "Hello, World!", AuthorName: "Sau Sheong"}
	post.Create()
	fmt.Println(post)
}
