package main

import "fmt"

func main() {
num := [...] int {1,2,3,4,5}
num2 := num
num[3] = 22
fmt.Println(num)
fmt.Println(num2)
}