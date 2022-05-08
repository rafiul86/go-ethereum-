package main

import (
	"fmt"
)

func main() {
	nyNum := printNum(12, 14)
	fmt.Println(nyNum)
}

func printNum (num1 int, num2 int) (result int) {
	result = num1 + num2
	return
}