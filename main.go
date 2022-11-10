package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five Of Diamonds"
	// fmt.Println("card:", card)

	cards := newDeck()
	cards = append(cards, "Six of Spades")

	fmt.Println("All cards:")
	fmt.Println("=========================")
	fmt.Println(cards.toString())

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Card on hands:")
	fmt.Println("=========================")

	fmt.Println(hand.toString())

	fmt.Println("Remaning cards:")
	fmt.Println("=========================")
	fmt.Println(remainingCards.toString())

	// messages := remainingCards.toString()
	// fmt.Println([]byte(messages))

}
