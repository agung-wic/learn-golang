package main

import (
	"fmt"
)

func main() {
	// if the constant is not used then it does not show an error
	// constant cannot be declared using ":" like a variable
	const name string = "Agung"
	const city = "Surabaya"
	const age = 23

	const (
		firstName = "Agung"
		lastName  = "Wicaksono"
	)

	fmt.Println(name)
	fmt.Println(city)
	fmt.Println(age)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
