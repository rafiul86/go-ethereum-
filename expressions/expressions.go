package expressions

func Add(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum = sum + value
	}
	return sum
}