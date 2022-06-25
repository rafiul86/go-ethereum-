package main

import (
	"expressions/expressions"
	"fmt"
)

func main() {
	num := []float64{1.08, 2.06, 3.06, 4.07, 5.90}
	fmt.Printf("%f\n", expressions.Add(num...))
}

