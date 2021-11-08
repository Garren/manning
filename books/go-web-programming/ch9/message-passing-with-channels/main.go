package main

import (
	"fmt"
	"time"
)

// throw a number in a sequence to the passed in channel
func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i // write to the channel
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c // read from the channel
		fmt.Println("Caught <<", num)
	}
}

func callerA(c chan string) {
	c <- "Hello, World!"
}
func callerB(c chan string) {
	c <- "Hola Mundo!"
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
	{
		fmt.Println("unbuffered")
		// use an unbuffered channel
		// this is synchronous and the values will be displayed in order
		c := make(chan int)
		go thrower(c)
		go catcher(c)
		time.Sleep(100 * time.Millisecond)
	}

	{
		fmt.Println("buffered")
		// use a buffered channel.
		// this is asynchronous and won't block so long as something is in the
		// channel
		c := make(chan int, 3)
		go thrower(c)
		go catcher(c)
		time.Sleep(100 * time.Millisecond)
	}
	{
		fmt.Println("select")
		// make a pair of unbuffered channels of string
		a, b := make(chan string), make(chan string)

		// each routine will write a string to their channel
		// five times
		go callerA(a)
		go callerB(b)

		for i := 0; i < 5; i++ {
			// give the goroutines a bit to startup
			time.Sleep(1 * time.Microsecond)

			select {
			case msg := <-a:
				fmt.Printf("%s from A\n", msg)
			case msg := <-b:
				fmt.Printf("%s from B\n", msg)
			default:
				fmt.Println("Default")
			}
		}
	}
	{
		fmt.Println("select with close")
		// make a pair of unbuffered channels of string
		a, b := make(chan string), make(chan string)

		go func(c chan string) {
			c <- "hello world!"
			close(c)
		}(a)

		go func(c chan string) {
			c <- "hola mundo!"
			close(c)
		}(b)

		var msg string
		ok1, ok2 := true, true

		// as soon as the func writes its string it closes the channel
		// we use a multivalue return to read the value from the channel,
		// and to determine if it's closed.
		for ok1 || ok2 {
			select {
			case msg, ok1 = <-a:
				if ok1 {
					fmt.Printf("%s from A\n", msg)
				}
			case msg, ok2 = <-b:
				if ok2 {
					fmt.Printf("%s from B\n", msg)
				}
			}
		}
	}
}
