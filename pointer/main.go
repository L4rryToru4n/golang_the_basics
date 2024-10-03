package main

import "fmt"

// By default variables in Golang are passing by values not by reference
// This means that all variables that are sent into functions, methods or
// other variables are their duplicated values (copied), not their references
// This also means it is safe to change the newly assigned variable's
// data with the previous variable since it will only change the current
// variable's data and will not affect the previously assigned variable's
// data 

type Address struct {
	Street string
	City string
	State string
	Country string
}

func main() {

	address_1 := Address{"Steinway St", "New York City", "New York", "USA"}
	address_2 := address_1

	address_2.Street = "Joyce St"
	address_2.City = "Austin"
	address_2.State = "Texas"

	fmt.Println(address_1) // address_1 values not changing
	fmt.Println(address_2)
	fmt.Println("--------------------------------------")

	// To make it pass by reference, add & to the variable assignment
	
	var newAddress_1 Address = Address{"Steinway St", "New York City", "New York", "USA"}
	var newAddress_2 *Address = &newAddress_1 // pointer

	newAddress_2.Street = "Joyce St"
	newAddress_2.City = "Austin"
	newAddress_2.State = "Texas"

	fmt.Println(newAddress_1) // newAddress_1 values changed
	fmt.Println(*newAddress_2)

	var newAddress_1_Ref *Address = &newAddress_1

	fmt.Println(newAddress_1_Ref) // memory address of newAddress_1
	fmt.Println(newAddress_2) // memory address of newAddress_2
}