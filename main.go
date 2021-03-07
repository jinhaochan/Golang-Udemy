package main

import "fmt"

func main() {
	// equivalent to `var card string = "Ace of Spades"`
	card := "Ace of Spades"

	// Golang being a static language, you cannot assign another type to it
	// such as `card = 1`

	// reassignment of value to card
	// you cannot run `card := "Five of Diamonds"` again
	card = "Five of Diamonds"

	fmt.Println(card)
}
