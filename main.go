package main


func main(){
	card := cards{"Ace of Spades", "Two of Diamonds", "Three of Hearts"}
	card = append(card, "Four of Clubs")
	card.print()
}