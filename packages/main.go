package main

import (
	"fmt"
	MyPackage "golang_basics/packages/my_package"
)

// block -> package -> universe

var myGlobal = 5

func main() {
	{
		var myInteger = 2
		fmt.Println(myInteger)
	}

	var myInteger = 1

	fmt.Println(myInteger)
	fmt.Println(myGlobal)

	nonMain()

	result_int := MyPackage.MyNewInteger
	result_func := MyPackage.SayHello("You")

	fmt.Println(result_int)
	fmt.Println(result_func)
}

func nonMain() {
	fmt.Println(myGlobal)
}
