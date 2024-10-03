package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Country string
}

// Making a new pointer

func main() {
	address_1 := new(Address)

	address_2 := address_1

	address_2.Country = "USA"

	fmt.Println(address_1) // memory address changed
	fmt.Println(address_2)
}
