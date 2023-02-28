package main

import "fmt"

func main() {
	type Address string
	type Age uint8

	var myAddress Address = "Madiun"
	var myAge Age = 23

	fmt.Println(myAddress)
	fmt.Println(myAge)
}
