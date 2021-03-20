package main

import (
	"fmt"
	"io"
	"os"
)

func isErr(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1]

	// file implements a Reader interface
	file, err := os.Open(args)

	isErr(err)

	io.Copy(os.Stdout, file)
}
