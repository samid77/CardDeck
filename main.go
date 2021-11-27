package main

func main() {

	// hand, remaningCards := deal(cards, 5)

	// hand.print()
	// remaningCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("firstCards")

	// cards := newDeckFromFile("firstCards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
