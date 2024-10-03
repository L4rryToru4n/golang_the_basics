package main

import "fmt"

func main() {
	counter := 0
	
	// Anonymous function
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}