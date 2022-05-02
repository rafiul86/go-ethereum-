package main

import "fmt"

func main() {
num := [] int {1,2,3,4,5}
num2 := [] int {6,7,8,9,10}
num3 := append(num, num2...)
num[3] = 22
fmt.Println(num)
fmt.Println(num2)
fmt.Println(num3)
}