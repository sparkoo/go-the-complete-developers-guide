package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	remainingDeck.print()
	hand.print()

	fmt.Println(hand.toString())
	hand.saveToFile("hand_deck")
}
