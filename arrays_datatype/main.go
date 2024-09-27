package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Lee";
	names[1] = "'Husker'";
	names[2] = "Adama";

	for i := 0; i < 3; i++ {
		fmt.Printf(names[i] + " ");
	}
	fmt.Printf("\n");

	var values = [3]int{
		90,
		80,
	}

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i]);
	}

	var emptyValues = [3]int{}

	for i := 0; i < len(emptyValues); i++ {
		fmt.Println(emptyValues[i]);
	}

	// For automatically count the length

	var autoVal = [...]int {
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(len(autoVal));

}
