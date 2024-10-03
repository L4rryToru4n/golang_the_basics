package main

import "fmt"

func random() interface{} {
	return "OK"
}

func newRandom() any {
	// return string
	// return "OK"

	// return integer
	return 10

	// return boolean
	// return true
}

func main() {
	var result any = random()
	var resultStr string = result.(string)

	fmt.Println(resultStr)

	var resultInt int = result.(int) // <- panic, cannot do conversion because it is an invalid datatype (string)
	fmt.Println(resultInt)

	// Correct way to do type assertions,
	// check the datatype first
	
	var newResult any = newRandom()

	switch value := newResult.(type) {
		case string:
			fmt.Println("String", value)
		case int:
			fmt.Println("Integer", value)
		default:
			fmt.Println("Unknown")
	}
}