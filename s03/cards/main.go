package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of diamonds"}
	fmt.Println(cards)

	cards2 := append(cards, "Six of spades")
	fmt.Println(cards2)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for card := range cards2 {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Ace of spades"
}
