package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func main() {
	dsn := "host=localhost port=5432 user=postgres dbname=chitchat sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	Check(err)
	defer db.Close()

	rows, err := db.Query("select id, uuid, name, email, password, created_at from users")
	if err != nil {
		Check(err)
	}
	users := []User{}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()

	fmt.Println(users)
}
