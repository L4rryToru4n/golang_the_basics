package CoreFunctions

import "fmt"

type InvalidIDError struct {
	Message string
}

type NotFoundError struct {
	Message    string
	CodeNumber string
}

func (i *InvalidIDError) Error() string {
	return i.Message
}

func (n *NotFoundError) Error() string {
	return n.Message + n.CodeNumber
}

func Recovery() {

	message := recover()
	if message != nil {
		fmt.Println("An error occured:", message)
	}
}

func InvalidRecovery() {
	message := recover()
	invalidErr := &InvalidIDError{"Invalid value."}
	fmt.Println(message, invalidErr)
}

func NotFoundRecovery() {

	message := recover()
	notFoundErr := &NotFoundError{"The ID number was not found ", "[CE-001]"}
	fmt.Println(message, notFoundErr)
}
