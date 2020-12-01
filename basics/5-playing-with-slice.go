package main

import "fmt"

func main() {
	cards := []string{
		newCard(), newCard(), newCard(),
	}
	cards = append(cards, "Other card")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
