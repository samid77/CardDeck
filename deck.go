package main

import "fmt"

type deck []string

//Create new deck of cards and return list of cards/deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Log out all contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
