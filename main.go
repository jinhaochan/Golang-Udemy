package main

func main() {
	myDeck := newDeck()

	hand, remainingCards := deal(4, myDeck)

	remainingCards.print()

	hand.print()

}
