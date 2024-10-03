package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (n *NotFoundError) Error() string {
	return n.Message
}


// Error is an interface so must be returned
// in a pointer datatype

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Validation error."}
	} else if id != "123" {
		return &NotFoundError{"Data not found."}
	} else {
		return nil
	}
}

func main() {
	err := SaveData("123", nil)

	if err != nil {
		// This will lead to an error

		// type conversion into pointer
		if ValidationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Validation error: ", ValidationErr.Error())
		} else if NotFoundErr, ok := err.(*NotFoundError); ok {
			fmt.Println("Not found error: ", NotFoundErr.Error())
		} else {
			fmt.Println("Unknown error: ", err.Error())
		}
	} else {
		// success
		fmt.Println("success")
	}
}