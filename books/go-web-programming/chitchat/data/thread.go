package data

import (
	"fmt"
	"time"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// view helper to format post dates
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get a usr for a post
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("select id, uuid, name, email, created_at from users where id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// get all threads
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order by created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

// find a thread by its uuid
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("select id, uuid, topic, user_id, created_at from threads where uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user for a thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("select id, uuid, name, email. password, created_at from users where id = $1", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// view helper for date formatting in threads
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts for a given thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get all threads for the current post
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = $1", thread.Id)
	if err != nil {
		err = fmt.Errorf("Posts/Query: %w", err)
		return
	}

	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			err = fmt.Errorf("Posts/Scan: %w", err)
			return
		}
		posts = append(posts, post)
	}

	rows.Close()
	return
}

// Create a thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("CreateThread/Prepare: %w", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	if err != nil {
		err = fmt.Errorf("CreateThread/QueryRow: %w", err)
	}
	return
}

// Create a post for a given thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("CreatePost/Prepare: %w", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	if err != nil {
		err = fmt.Errorf("CreatePost/QueryRow: %w", err)
	}
	return
}
