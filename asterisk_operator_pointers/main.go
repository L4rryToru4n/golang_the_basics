package main

import "fmt"

type Address struct {
	Street string
	City string
	State string
	Country string
}

func main() {
	var address_1 Address = Address{"Steinway St", "New York City", "New York", "USA"}
	var address_2 *Address = &address_1

	address_1.Street = "Joyce St"
	address_1.City = "Austin"
	address_1.State = "Texas"

	address_2 = &Address{"Steinway St", "New York City", "New York", "USA"}
	
	fmt.Println(address_1) // address_1 values not changing
	fmt.Println(address_2)
	fmt.Println("--------------------------------------")

	// If we want the reference of address_2 to be affecting address_1
	// too then we use the asterisk operator to get the reference of
	// address_3 from address_4
	var address_3 Address = Address{"Steinway St", "New York City", "New York", "USA"}
	var address_4 *Address = &address_3
	
	*address_4 = Address{"Broadwalk St", "Moon", "Sol", "Milky Way"}
	
	fmt.Println(address_3)
	fmt.Println(address_4)

}