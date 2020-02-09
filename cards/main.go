package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
