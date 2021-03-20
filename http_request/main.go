package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func isErr(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}

type logWriter struct{}

func main() {

	// creating a byte slice with 9999 empty slots
	// we need to do this because the reader function does not automatically resize a slice
	// it will keep reading data until the slice is full
	// by declaring a slice []byte{}, the size would be 0, and no data would be read
	body := make([]byte, 9999)

	// have to append "http"
	res, err := http.Get("http://www.google.com")

	isErr(err)

	// we dont have to implement the Read function
	// its already implemented for us
	// takes in a byte slice
	// the reader pushes data into the byte slice
	// returns int and error
	// int = number of bytes that was read into the slice
	// error = error if any, if not nilc
	res.Body.Read(body)

	fmt.Println(string(body))

	lw := logWriter{}

	// quick hack
	// copies the body data into stdout
	// io.Copy takes in a Writer and a Reader
	io.Copy(lw, res.Body)

}

// creating a custom writer
// logWriter implements Writer interface
func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))

	return len(b), nil
}
