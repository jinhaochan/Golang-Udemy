package main

import "fmt"

// Create new type of deck
// which is a slice of string
// because there is no number in the square brackets, its considered a slice
// an array looks like this
// var myArray [256]int
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

// do not need to add a receiver. this is a case by case basis
// experimenting with empty returns
func newDeck() (cards deck) {
	//cards := deck{}

	cardSuites := []string{"Clubs", "Spades", "Hearts", "Diamonds"}

	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King", "Ace"}

	// use underscores for variables not used
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return
}
