package main

import "fmt"

func main() {
	cards := []string{newCard(), "Queen of Hearts"}
	cards = append(cards, "King of Diamonds")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
