package main

import "fmt"

func main() {
	ok := true
	for i := 0; i < 6; i++ {
		if i%4 == 3 {
			ok = false
			switch ok {
			case true:
				fmt.Println(i, "ok is true")
			case false:
				fmt.Println(i, "ok is false")
			default:
				fmt.Println(i, "ok is unknown")
			}
		}
	}
}
