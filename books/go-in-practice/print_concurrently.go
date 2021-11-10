package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {

	// count called as a goroutine
	go count()

	// sleep the main thread for two seconds, this give our count() 2ms to print
	// to the console.
	time.Sleep(time.Millisecond * 2)

	// after 2ms we print howdy.
	fmt.Println("Hello, world!")

	// sleep for another 5ms giving our count() the ability to finish the last
	// 3 console writes
	time.Sleep(time.Millisecond * 5)
}
