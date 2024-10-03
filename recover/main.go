package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func recovery() {
	message := recover()
	fmt.Println("a panic function was called: ", message)
}

func runApp(error bool) {

	defer recovery()
	defer endApp()

	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("App continue..")
}