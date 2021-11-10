package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user  // "embedded" type - no instance name
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "Jon Smith",
			email: "Jon@example.com",
		},
		level: "root",
	}

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
