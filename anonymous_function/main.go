package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "blocked")
	} else {
		fmt.Println("Hello", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Cylon"
	}

	registerUser("Cylon", blacklist)

	registerUser("Lee", func(name string) bool {
		return name == "Cyclon"
	})
}
