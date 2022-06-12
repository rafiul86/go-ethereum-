package main

import (
	"app/doctor"
	"fmt"
)


func main() {

	nur := doctor.Nurse{Name: "Larry",Age: 30}

	age := nur.Talk()
	fmt.Println(age)
}
	
