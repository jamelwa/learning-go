package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// create a function with return type
func newCard()  string {
	return "Ace of Spades"
}