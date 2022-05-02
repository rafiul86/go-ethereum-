package main

import "fmt"
type people struct {
	count string
	time int
}
type person struct {
	people
	name string
	age int
}
func main() {
Loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j >= 15 { break Loop }
			fmt.Println(i, "*", j, "=", i*j)
		}
		}
}