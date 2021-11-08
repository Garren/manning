package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// create a session for a user
func (user *User) CreateSession() (session Session, err error) {
	statement :=
		"insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("CreateSession()/Prepare: %w", err)
		return
	}
	defer stmt.Close()
	err = stmt.
		QueryRow(createUUID(), user.Email, user.Id, time.Now()).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		err = fmt.Errorf("CreateSession(): %w", err)
	}
	return
}

// get a session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.
		QueryRow("select id, uuid, email, user_id, created_at from sessions where user_id = $1", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		err = fmt.Errorf("Session(): %w", err)
	}
	return
}

// create a user and persist to database
func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("Create()/Prepare: %w", err)
		return
	}
	defer stmt.Close()

	err = stmt.
		QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	if err != nil {
		err = fmt.Errorf("Create(): %w", err)
	}
	return
}

// delete a user
func (user *User) Delete() (err error) {
	statement := "delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("Delete(): %w", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return fmt.Errorf("Delete(): %w", err)
	}
	return
}

// update a user's information and persist to db
func (user *User) Update() (err error) {
	statement := "update users set name = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("Update(): %w", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("Update: %w", err)
	}
	return
}

// delete all users
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// get all users
func Users() (users []User, err error) {
	rows, err := Db.Query("select id, uuid, name, email, password, created_at from users")
	if err != nil {
		err = fmt.Errorf("Users(): %w", err)
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			err = fmt.Errorf("Users()/rows.Next(): %w", err)
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// find a single user by their email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id, uuid, name, email, password, created_at from users where email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		err = fmt.Errorf("UserByEmail: %w", err)
	}
	return
}

// find a single user by their uuid
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id, uuid, name, email, password, created_at from users where uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		err = fmt.Errorf("UserByUUID: %w", err)
	}
	return
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// see if a session exists in the db
func (session *Session) Check() (valid bool, err error) {
	err = Db.
		QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		err = fmt.Errorf("Check(): %w", err)
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// Get the user's session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id, uuid, name, email, created_at from users where id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		err = fmt.Errorf("User(): %w", err)
	}
	return
}

// Delete all sessions
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}

// delete a session by its uuid
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return
}
