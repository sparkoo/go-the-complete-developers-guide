package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	remainingDeck.print()
	hand.print()
}
