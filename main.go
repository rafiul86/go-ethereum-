package main

import "fmt"

func main() {
	greetings := "Hello"
	greetings = (greetings + " World")
	fmt.Printf("%v %T", greetings, greetings)
}