package main

import "fmt"

func main() {
	var value32 int32 = 250000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)

	name := "Agung"
	char := name[3]
	charString := string(char)

	fmt.Println(charString)
}
