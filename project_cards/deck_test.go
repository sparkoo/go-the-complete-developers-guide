package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected size of deck to be 52, but got [%v]", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card to be Ace of Diamonds, but got [%v]", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got [%v]", d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"

	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if !reflect.DeepEqual(d, loadedDeck) {
		t.Errorf("Expected new and loaded decks to be same")
	}

	os.Remove(testFile)
}
