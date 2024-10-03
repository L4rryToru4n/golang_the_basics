package main

import (
	"fmt"
	"errors"
)

// Package error is an interface
/*
	 type error interface {
	 	Error() string
	 }
*/

func Division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Cannot be divided by 0")
	} else {
		return value / divider, nil
	}
}

func main() {
	fmt.Println(Division(4, 2))
	fmt.Println(Division(4, 0)) // <- Will return an error

	result, err := Division(100, 10)
	// result, err := Division(100, 0)

	if err == nil {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("ERROR: ", err.Error())
	}

}