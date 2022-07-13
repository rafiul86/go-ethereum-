package main

import (
	"fmt"
	"sync"
)

func printSlice(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
	
}


func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	mySlice := []string{
		"go",
		"solidity",
		"python",
		"javascript",
	}
	for _, v := range mySlice {
		go printSlice(v, &wg)
	}
	wg.Wait()
}