package main

import (
	"app/doctor"
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	sayHello := doctor.Intro()
	fmt.Println(sayHello)
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}
