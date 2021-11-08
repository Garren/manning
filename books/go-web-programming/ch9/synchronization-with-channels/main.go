package main

import (
	"fmt"
	"time"
)

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func main() {
	/*
		make an unbuffered channel.
		unbuffered channels are syncrhonous - readers will block when there's
		nothing to read. writers will block if there's a value that hasn't been
		read.
		ch := make(chan int)

		make a buffered channel of 10 integers
		ch := make(chan int, 10)

		put a value into a channel
		ch <- 1

		read a value out of a channel
		i := <-ch

		by default channels are bi-directional.

		make a send only channel of string
		ch := make(chan <- string)

		make a receive only channel of string
		ch := make(<-chan string)
	*/

	w1, w2 := make(chan bool), make(chan bool)

	go printNumbers2(w1)
	go printLetters2(w2)

	// try to read from the channels as soon as you fire the goroutines above.
	// these will block until the routines above complete and write to their
	// respective channels, at which point the reads below will unblock.
	<-w1
	<-w2
}
