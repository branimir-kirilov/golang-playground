package main

import (
	"fmt"
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
