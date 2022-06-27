package main

import (
	"fmt"
)

func main() {
	courses := []string{"golang", "python", "solidity"}

	for _, course := range courses {  
		fmt.Println(course)
	}
}
