package main

import (
	"fmt"
	"io/ioutil"
)

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

	cards.saveToFile("cards.txt")

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Card on hands:")
	fmt.Println("=========================")

	fmt.Println(hand.toString())
	hand.saveToFile("hands.txt")

	fmt.Println("Remaning cards:")
	fmt.Println("=========================")
	fmt.Println(remainingCards.toString())

	// messages := remainingCards.toString()
	// fmt.Println([]byte(messages))

	fmt.Println("===================================")

	remainingCards.saveToFile("remainingCard.txt")
	res, isOk := ioutil.ReadFile("cards.txt")

	if isOk == nil {
		fmt.Println("reading files...")
		fmt.Println(string(res))

	}

	fmt.Println("===================================")

	newDeck := newDeckFromFile("cadrds.txt")
	fmt.Println("new deck from file:", newDeck.toString())

}
