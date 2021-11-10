package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello"))

	fmt.Fprintf(&b, "World!\n") // bytes.Buffer implements io.Writer

	io.Copy(os.Stdout, &b) // bytes.Buffer implements io.Reader
}
