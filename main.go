package main

import (
	"app/doctor"
	"bufio"
	"fmt"
	"os"
)
func main() {
	sayHello := doctor.Intro()
	fmt.Println(sayHello)
	reader := bufio.NewReader(os.Stdin)
	for {
		userInput, _ := reader.ReadString('\n')
	fmt.Println(userInput)
	}
}
