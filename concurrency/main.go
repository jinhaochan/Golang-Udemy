package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{"http://www.google.com",
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com"}

	// making a channel of strings
	c := make(chan string)

	// iterate the sites and query them sequentially
	for _, val := range sites {
		// this is a blocking call, and we should make this run concurrently
		// checkLink(val)

		// creating a new go routine. akin to multithreading
		// passing in a channel so the routines can talk to each other
		go checkLink(val, c)
	}

	// waiting for the response for each routine that was started
	// this for loop would not run forever because waiting for the channel is a blocking operation
	// range c is a blocking, because it iterates over a channel
	for l := range c {
		// function literal, similar to lambda in python
		// we need to pass in l because l is define outside the scope of the function literal
		// for the link to be updated within the function literal, we need to pass it in
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Site", link, "is down!")
		c <- link
	} else {
		fmt.Println("Site", link, "is up!")
		c <- link
	}
}
