package main

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suits+" of "+values+" ,")
		}
	}
	return cards
}