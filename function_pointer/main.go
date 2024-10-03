package main

import "fmt"

type Address struct {
	Street string
	City string
	Region string
	Country string
}

// By default function parameters in Golang are pass by value, that means
// the value would be copied then sent into the function.
// To be able to change the original value, pointers can be passed to
// functions as parameters.
func ChangeAddressInfo(address *Address) {
	address.Street = "David St"
	address.City = "Leeds"
	address.Region = "Yorkshire and Humber"
	address.Country = "United Kingdom"
}

func main() {

	var address_one *Address = &Address{}
	var address_two Address = Address{}

	ChangeAddressInfo(address_one)
	ChangeAddressInfo(&address_two)

	fmt.Println(address_one)
	fmt.Println(address_two)
}