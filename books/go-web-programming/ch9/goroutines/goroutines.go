package main

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func printNumbers3() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d ", i)
	}
}

func printLetters3() {
	for i := 'A'; i < 'A'+10; i++ {
		//fmt.Printf("%c ", i)
	}
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func print2() {
	printNumbers2()
	printLetters2()
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}

func print3() {
	printNumbers3()
	printLetters3()
}

func goPrint3() {
	go printNumbers3()
	go printLetters3()
}

func main() {
}
