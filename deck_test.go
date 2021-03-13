package main

import (
	"os"
	"testing"
)

// captital T for Test, and followed by the function we want to test
func TestNewDeck(t *testing.T) {
	// checking the number of cards in the deck
	// checking if each card is unique

	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}

	if d[0] != "One of Clubs" {
		t.Errorf("Expected first card to be One of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Expected last card to be Ace of Diamonds, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
