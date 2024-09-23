package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello everyone !")

	// Full Declaration
	// var name type: value
	var myNumber int = 1

	var myGreeting string = "Helloo"

	// Short Declaration
	// name := expression
	myBoolean := false

	fmt.Println(myNumber, myGreeting)
	fmt.Println(myBoolean)

	// POINTERS
	var x int = 1
	// variable with a type of pointer integer and value of a memory address
	var pointer_x *int = &x

	// print the memory address and the value
	fmt.Println(pointer_x, *pointer_x)

	// TYPE DECLARATIONS
	// fahrenheit & celcius
	type fahrenheit int
	type celcius int

	var f fahrenheit = 32
	var c celcius = 0

	// The purpose of declaring a type is to not able to re-assigned a variable
	//  into another type. It will be incompatible types
	// c = f
	
	// casting a declared type by converting fahrenheit to celcius
	c = celcius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}