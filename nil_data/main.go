package main

import "fmt"

// func NilName(name string) string {
// 	if name == "" {
// 		return nil // <- Cannot be used because not in the allowed datatypes
// 	} else {
// 		return name
// 	}
// }

func NilMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string] string {
			"name": name,
		}
	}
}

func main() {

	// Nil represents data with an empty value
	// Can only be used in datatypes of
	// interface, functions, maps, pointers
	// and channels

	data := NilMap("")
	// data := NilMap("Kara Thrace")
	
	if data == nil {
		fmt.Println("Empty data")
	} else {
		fmt.Println(data["name"])
	}
}
