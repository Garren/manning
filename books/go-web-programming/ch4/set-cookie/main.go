package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co.",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	c3 := http.Cookie{
		Name:     "third_cookie",
		Value:    "Learning Go",
		HttpOnly: true,
	}
	c4 := http.Cookie{
		Name:     "fourth_cookie",
		Value:    "O'Reilly",
		HttpOnly: true,
	}
	http.SetCookie(w, &c3)
	http.SetCookie(w, &c4)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("second_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get second cookie")
	}

	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)

	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/set-cookie", setCookie)
	http.HandleFunc("/get-cookie", getCookie)

	server.ListenAndServe()
}
