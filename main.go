package main

import (
	"fmt"
	"strconv"
)
func main() {
	i, err := strconv.Atoi("65a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

