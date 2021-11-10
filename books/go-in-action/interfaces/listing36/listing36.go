package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func main() {
	u := user{"Bill", "bill@example.com"}
	sendNotification(u)

	//$ go run listing36.go
	//# command-line-arguments
	//./listing36.go:21:18: cannot use u (type
	//	user) as type notifier in argument to
	// sendNotification:
	//        user does not implement notifier
	// (notify method has pointer receiver)

	// call fails because we've implemented
	// notify using a pointer reciever and we
	// called with a value.
}

func sendNotification(n notifier) {
	n.notify()
}
