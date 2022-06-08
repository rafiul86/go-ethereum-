package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press Enter when ready."
// main function of Guess number game
func main() {
	// declare essential variables

	var firstNumber = 2
    var secondNumber = 5
	var subtraction = 7
	var answer int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Guess Number Game!")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')
	fmt.Println("Multiply your nymber by", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Divide the result the number by you originally thought of",  prompt)
	reader.ReadString('\n')
	fmt.Println("Now Subtract", subtraction, prompt)
	reader.ReadString('\n')
	answer = firstNumber * secondNumber - subtraction
	fmt.Println("Your answer is", answer)
}
