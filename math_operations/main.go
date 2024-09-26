package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	var d = a * b

	fmt.Println(d)

	var e = d % a

	fmt.Println(e)

	// Augmented Assignments
	fmt.Println("-------")
	var varOne int

	varOne += 10
	fmt.Println(varOne)

	varOne -= 5
	fmt.Println(varOne)

	varOne *= 2
	fmt.Println(varOne)

	varOne /= 2
	fmt.Println(varOne)

	varOne %= 5
	fmt.Println(varOne)

	// Unary operators
	fmt.Println("-------")
	var varTwo = 1

	varTwo++
	fmt.Println(varTwo)

	varTwo--
	fmt.Println(varTwo)

	fmt.Println(-varTwo)

	var varThree = true

	fmt.Println(!varThree)

}
