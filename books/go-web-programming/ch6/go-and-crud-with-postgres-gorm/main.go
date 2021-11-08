package main

import (
	"fmt"

	// this package has its own init() and will register the driver itself
	// on import. Otherwise we'd have to call
	//  sql.Register("postgres", &drv{})
	// the _ indicates that we're not to use the driver explicitly, instead we
	// rely on the database/sql interfaces it implements
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	// go get github.com/lib/pq
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string    `sql: "not null"`
	Comments  []Comment // note: slices are actually pointers
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql: "not null"`
	PostId    int    `sql: "index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=mypass sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello, World!", Author: "Sau Sheong"}
	Db.Create(&post)

	comment := Comment{Content: "Good point!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
