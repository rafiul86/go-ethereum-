package main

import (
	"fmt"
)
 
func agey() {
	fmt.Println("Agey")
}
func pore() {
	fmt.Println("pore")
}

func ahare() {
	fmt.Println("ahare")
}

func main() {
	defer agey()
	pore()
	defer ahare()
}