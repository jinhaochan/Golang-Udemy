package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}

	// append does not modify the actual slice, but returns a new slice
	cards = append(cards, "Six of Spades")

	myNewDeck := newDeck()

	myNewDeck.print()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
