package main

import "fmt"

func main() {

	/*
		If else expression
	*/

	// var value = 0
	var value = 1
	// var value = 2

	if value == 0 {
		fmt.Println("True, it's zero")
	} else if value == 1 {
		fmt.Println("True, it's one")
	} else {
		fmt.Println("False, it isn't one or zero")
	}

	// if short statement. Used if the variable
	// is not going to be used anywhere

	// var name = "Kara 'Starbuck' Thrace"
	var name = "Starbuck"

	if length := len(name); length > 8 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Length OK")
	}

	/*
		Switch expresion
	*/
	// var key = "Esc"
	var key = "2"

	switch key {
		case "1":
			fmt.Println("Power Management Menu")
		case "2":
			fmt.Println("Infantry Training Menu")
		case "3":
			fmt.Println("Spacecraft Production Menu")
		default:
			fmt.Println("Exit Management Menu")
	}

	// switch short statement. Used if the variable
	// is not going to be used anywhere

	// var alias = "Apollo"
	var alias = "Lee 'Apollo' Adama"

	switch length := len(alias); length > 8 {
		case true:
			fmt.Println("Alias is too long")
		case false:
			fmt.Println("Length OK")
	}

	// switch without condition
	
	var length = len(alias)
	switch {
		case length > 10:
			fmt.Println("Alias is far too long")
		case length > 8:
			fmt.Println("Alias is too long")
		default:
			fmt.Println("Length OK")
	}
}
