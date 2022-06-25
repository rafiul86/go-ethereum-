package main

import "fmt"

func main() {
	fmt.Printf("%f\n", add(1.08, 2.06, 3.06, 4.07, 5.90))
}

func add(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum = sum + value
	}
	return sum
}