package main

import "fmt"

// Create new type of deck
// which is a slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
