package main

import "fmt"

func main() {

	// map[key_datatype] value_datatype

	// var person = map[string] string {
	//	"name": "Raider",
	//	"rank": "Captain",
	//}

	var person = map[string]string{}

	person["name"] = "Raider"
	person["rank"] = "Captain"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["rank"])

	fmt.Println(person["noKey"]) // <- If there's no key then will access the default value

	/* 
		Map function
	*/

	// make(map[key_datatype]value_datatype)
	var book = make(map[string]string)
	book["title"] = "How to Win negotiation"
	book["author"] = "Bob Ferguson"
	book["not_title"] = "0"

	fmt.Println(book)

	delete(book, "not_title")

	fmt.Println(book)

	// len()
	fmt.Printf("Book length: %d\n", len(book))
}
