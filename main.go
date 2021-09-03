package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	cards.shuffle()
	// cards.print()
}
