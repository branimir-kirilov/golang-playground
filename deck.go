package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of Deck which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function for type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func loadDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading file", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}
