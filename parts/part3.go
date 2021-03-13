package main

import (
	"fmt"
)

func main() {
	myDeck := newDeck()

	hand, remainingCards := deal(4, myDeck)

	fmt.Println(hand)

	remainingCards.saveToFile("myfile.txt")

	deckFromFile := newDeckFromFile("myfile.txt")

	fmt.Println(deckFromFile)

}
