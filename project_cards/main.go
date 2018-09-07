package main

func main() {
	deckFile := "my_deck.data"

	cards := newDeck()

	hand, _ := deal(cards, 5)

	//remainingDeck.print()
	//hand.print()

	hand.print()
	hand.saveToFile(deckFile)

	loadedDeck := newDeckFromFile(deckFile)
	loadedDeck.print()

	loadedDeck.shuffle()
	loadedDeck.print()
}
