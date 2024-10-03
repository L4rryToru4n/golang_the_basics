package main

import (
	"fmt"
	"golang_basics/access_modifier/helper"
)

func main() {
	message := helper.ShowMsg("Hello from the other side.")
	// myMsg := helper.printMessage("Hello from the other side.") // <- cannot be accessed
	application := helper.Application
	// ver := helper.version // <- cannot be accessed

	fmt.Println(message, application)
	// fmt.Println(myMsg, ver) // Error. Not exported by package helper
}
