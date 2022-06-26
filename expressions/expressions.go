package expressions

func Add(values ...float64) (sum float64) {
	for _, value := range values {
		sum = sum + value
	}
	return
}