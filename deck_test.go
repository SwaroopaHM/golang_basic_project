package main

import (
	"os"
	"testing"
)

//test handler 't'
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected the deck length to be 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be 'Ace of Spades', but got %v", d[0])
	}
	if d[len(d)-1] != "Four of King" {
		t.Errorf("Expected the first card to be 'Four of King', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingDeck")

	d := newDeck()
	d.saveToFile("_testingDeck")
	loadedDeck := newDeckFromFile("_testingDec")

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected the length of the loadedDeck to be 16, but got %v", len((loadedDeck)))
	}
	os.Remove("_testingDeck")
}
