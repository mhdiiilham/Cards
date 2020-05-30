package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected first deck length of 52, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got: %v", d[len(d)-1])
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()
	if d[0] == "Ace of Spades" {
		t.Errorf("Expected Ace of Spades is not in the first index")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected first deck length of 52, but got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}