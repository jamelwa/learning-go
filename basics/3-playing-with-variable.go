package main

import "fmt"

func main() {

	//string value
	var card string = "This is a string"

	//alternatively you can declare variable like this
	card2 := "This is another string"

	//changing value
	card2 = "Try to change value"

	//you cannot declare new value with the same variable
	//card2 := "Try to re-define value"

	//boolean value
	var isRed bool = true

	fmt.Println(card)
	fmt.Println(card2)
	fmt.Println(isRed)
}
