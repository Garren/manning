package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello, World!")

	// write and read using ioutil. WriteFile and ReadFile work with byte slices
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// use a file struct for writing and reading. Need to create the file
	// first, and remember to defer Close
	file1, _ := os.Create("data2")
	bytes, _ := file1.Write(data)
	fmt.Printf("\nWrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
