package main

import "fmt"

// declaration of variables can be done outside the functions
// they just cannot be assigned a value
// e.g. `var card2 string = "Queen of Hearts"`

var card2 string

func main() {
	// equivalent to `var card string = "Ace of Spades"`
	card := "Ace of Spades"

	// Golang being a static language, you cannot assign another type to it
	// such as `card = 1`

	// reassignment of value to card
	// you cannot run `card := "Five of Diamonds"` again
	card = "Five of Diamonds"

	card2 = "Queen of Hearts"

	card3 := newCard()

	fmt.Println(card, ",", card2, ",", card3)
}

// return type of function declared after the function name
func newCard() string {
	return "Five of Diamonds"
}
