package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	fmt.Println(len(d))
	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	f := "Ace of Spades"
	if d[0] != f {
		t.Errorf("Expected first card to be %v, but it was %v", f, d[0])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	h, r := deal(d, 5)

	if len(h) != 5 {
		t.Fail()
	}

	if len(r) != 47 {
		t.Fail()
	}
}

func TestWriteToFile(t *testing.T) {
	fn := "_deckTestsFile"
	os.Remove(fn)

	deck := newDeck()
	errWrite := deck.writeToFile(fn)

	if errWrite != nil {
		t.Error("Error while writing the file!")
	}

	bs, errRead := os.ReadFile(fn)

	if errRead != nil {
		t.Error("Error while reading a file from the disk!")
	}

	s := strings.Split(string(bs), ",")

	if len(s) != 52 {
		t.Errorf("Expected the length to be 52 but it was %v", len(s))
	}

	if s[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades but it was %v", s[0])
	}

	os.Remove(fn)
}

func TestLoadDeckFromFile(t *testing.T) {
	fn := "_deckTestsFile"
	os.Remove(fn)

	d := newDeck()
	errWrite := d.writeToFile(fn)

	if errWrite != nil {
		t.Error("Error while writing the file!")
	}

	deck := loadDeckFromFile(fn)
	if len(deck) != 52 {
		t.Errorf("Expected the length to be 52 but it was %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades but it was %v", deck[0])
	}

	os.Remove(fn)
}
