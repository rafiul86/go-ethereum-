package main

import (
	"fmt"
	"os"
)
 
func main() {
	var newFile *os.File 
	fmt.Println(newFile)

	newFile, err := os.Create("test.txt")
	os.WriteFile("test.txt", []byte("Hello World"), 0644)
	if err != nil {
		fmt.Println(err, newFile)
	}
}