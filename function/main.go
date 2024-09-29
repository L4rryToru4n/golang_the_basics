package main

import "fmt"


func sayHello() {
	fmt.Println("Hello Everyone !")
}

func getPersonel(name string, rank string) (string, string) {
	return rank, name
}

func ignoredReturn() (int, string) {
	return 1, "Sunday"
}

func getFullPersonel() (name, rank, role string) {
	// Named Return Values
	name = "Lee 'Apollo' Adama"
	rank = "Liutenant"

	return role, rank, name
}

func sumAll(numbers ...int) int {
	// Variadic Function
	total := 0

	for _ ,number := range numbers {
		total += number
	}
	return total
}

func main() {
	sayHello()
	fmt.Println(getPersonel("Lee Adama", "Ensign"))
	_, dayString := ignoredReturn()
	fmt.Println(dayString)

	role, rank, name := getFullPersonel()
	role = "Viper Pilot"
	fmt.Println(role, rank, name)

	fmt.Println(sumAll(5, 4, 3, 2, 1))

	// Slice parameter
	numbers := []int {6, 7, 8, 9, 10}
	fmt.Println(sumAll(numbers...))
}