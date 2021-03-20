package main

import "fmt"

// the interface bot defines that the implementor must have a function
// getGreeting
type bot interface {
	getGreeting(string) string
	getBotVersion() string
}

type englishBot struct {
	version string
}
type spanishBot struct {
	version string
}

func main() {
	eb := englishBot{version: "1AB"}
	sb := spanishBot{version: "3RR"}

	printGreeting(eb)
	printGreeting(sb)
	printVersion(eb)
	printVersion(sb)
}

// passing in the interface bot
// what ever is pass in implements this interface
// therefore, whatever passed in needs to implement getGreeting function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting("ola"))
}

func printVersion(b bot) {
	fmt.Println(b.getBotVersion())
}

func (englishBot) getGreeting(s string) string {
	return "Hello!"
}

func (spanishBot) getGreeting(s string) string {
	return "ola!"
}

func (e englishBot) getBotVersion() string {
	return e.version
}

func (e spanishBot) getBotVersion() string {
	return e.version
}
