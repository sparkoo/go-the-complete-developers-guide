package main

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 5)

	//remainingDeck.print()
	//hand.print()

	hand.print()
	hand.saveToFile("hand_deck")

	loadedDeck, _ := loadFromFile("hand_deck")
	loadedDeck.print()
}
