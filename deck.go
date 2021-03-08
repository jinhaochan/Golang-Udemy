package main

import "fmt"

// Create new type of deck
// which is a slice of string
type deck []string

// receiver function
// by declaring the receiver below,
// any variable of type deck gets access to this print method
// d is the actual copy of the data we're working with
// the receiver is traditionally named as the first letter of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
