// Package main provides ...
package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email to %s\n", u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// create a value user type
	bill := user{"Bill", "bill@example.com"}
	bill.notify()

	// create pointer to a user instance
	lisa := &user{"Lisa", "lisa@example.com"}
	lisa.notify()

	// use a value type to call a method that accepts
	// a pointer. The compiler will take care of the
	// conversion to a pointer.
	bill.changeEmail("bill@corp.com")
	bill.notify()

	lisa.changeEmail("lisa@micro.com")
	// Here too, the compiler will take care of the
	// details for us. lisa will be dereferenced
	// automatically so we can call methods with a
	// value reciever with a pointer without worry
	lisa.notify()
}
