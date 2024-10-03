package main

import "fmt"

func anything() interface{} {
	return "Hi There!"
	// return 1
	// return true
}

func main() {

	anything := anything()

	fmt.Println(anything)
}
