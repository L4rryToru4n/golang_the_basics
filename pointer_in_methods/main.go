package main

import "fmt"

type Address struct {
	Street  string
	City    string
	Region  string
	Country string
}

// Even if method can be connected with struct, the data inside struct
// that are accessed by method are pass by value. To be able to change
// the original data, use pointer methods
// It is recommended to use pointer methods as the default of creating
// methods for structs
func (address *Address) ChangeAddressInfo() {
	address.Street = "David St"
	address.City = "Leeds"
	address.Region = "Yorkshire and Humber"
	address.Country = "United Kingdom"
}

func main() {
	var address Address = Address{"St Andrews St", "Northampton", "East Midlands", "United Kingdom"}
	address.ChangeAddressInfo()

	fmt.Println(address)
}
