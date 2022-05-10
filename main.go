package main

// necessary imports for the program to run
import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// main function for the program to run in the terminal
func main() {
	// create a new scanner to read from the terminal
	reader := bufio.NewReader(os.Stdin)

	// create a new doctor object from the doctor package
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	// loop until the user enters "quit"
	for {
		fmt.Print("->")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		if userInput == "quit" {
			break
		}else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}
