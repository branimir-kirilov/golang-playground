package main

func main() {
	cards := loadDeckFromFile("deck.txt")
	cards.print()
	cards.shuffle()
	cards.print()
}
