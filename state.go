package main

import "fmt"


type cards []string

func(c cards) print() {
	for i, card := range c{
		fmt.Println(i, card)
	}
}