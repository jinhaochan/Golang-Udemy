package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}

	// append does not modify the actual slice, but returns a new slice
	cards = append(cards, "Six of Spades")

	// range iterates over a collection, returning the index and element at each iteration
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
