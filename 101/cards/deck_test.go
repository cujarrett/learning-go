package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Test that a deck is created with x number of cards
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Test that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Test to make sure that te last card is a Four of Clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	// Test saving a deck to a file and then reading it back from a file
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in the deck, got %v", len(loadedDeck))
	}
}
