package main

import (
	"fmt"
)
func main() {
	a, b := 6, 0

	if b!=0 && division(a, b) == 2 {
		fmt.Println("a/b = 2")
	} else {
		fmt.Println("a/b != 2")
	}
}


func division(a, b int) int {
	return a / b
}

