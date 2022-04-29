package main

import "fmt"

func main() {
	cards := []int{2, 4, 6}
	cards = append(cards, 8)
	for index,  card := range cards {
		fmt.Println(index, card)
	}
}