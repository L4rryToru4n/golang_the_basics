package main

import "fmt"

// Struct
type Customer struct {
	Name    string
	Address string
	Age     int
}

// Struct Method
func (customer Customer) printCustomers(msg string) {
	fmt.Printf("%s Name is %s, address is %s, age is %d\n", msg, customer.Name, customer.Address, customer.Age)
}

func main() {

	// Struct literals A
	customerB := Customer{
		Name:    "Kevin",
		Address: "Space",
		Age:     29,
	}

	// Struct literals B
	customerC := Customer{"Scott", "Space", 31}

	var customerA Customer
	customerA.Name = "Christoper" // -> default value empty string ("")
	customerA.Address = "Space"   // -> default value empty string ("")
	customerA.Age = 37            // -> default value 0

	fmt.Printf("Name: %s, Address: %s, Age: %d\n", customerA.Name, customerA.Address, customerA.Age)
	fmt.Printf("Name: %s, Address: %s, Age: %d\n", customerB.Name, customerB.Address, customerB.Age)
	fmt.Printf("Name: %s, Address: %s, Age: %d\n", customerC.Name, customerC.Address, customerC.Age)

	// Struct Method Calls
	customerD := Customer{"James", "Space", 28}
	customerD.printCustomers("This is customer D data=")
}
