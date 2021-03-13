package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Clubs", "Spades", "Hearts", "Diamonds"}

	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King", "Ace"}

	// use underscores for variables not used
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(num int, d deck) (deck, deck) {
	return d[:num], d[num:]
}

// Converts a slice of strings, deck, into a single string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)

	// if there indeed is an error, we can either log the error and return a new deck
	// or log the error and quit the program

	if err != nil {
		fmt.Println("Error: ", err)
		// passing zero into Exit means that no errors were encountered
		// passing in a non-zero indicates that an error was encountered
		os.Exit(-1)
	}

	// casting the binary to a string
	s := string(bs)

	// splitting the string on comma
	d := strings.Split(s, ",")

	dd := deck(d)

	return dd
}

// we dont have to return a new deck, because this works on the value that was passed in
func (d deck) shuffle() {

	// creating a new rng seeded with nano time
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
