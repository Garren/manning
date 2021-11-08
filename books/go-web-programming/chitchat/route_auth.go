package main

import (
	"net/http"

	"git.sr.ht/~garren/go-web-programming/chitchat/data"
)

// GET /login
// display the login page
func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

// GET /signup
// display the signup page
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

// GET /authenticate
// authenticate a user given an email and password
func authenticate(w http.ResponseWriter, r *http.Request) {
	info("authenticate")
	err := r.ParseForm()
	if err != nil {
		danger(err, "Failed to parse form")
	}
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	info("found user", user)
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// GET logout
// log a user out using a cookie's uuid
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie { // sentinel error
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
