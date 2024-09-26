package main

import "fmt";

func main() {
	var integer32 int32 = 32768;
	var integer64 int64 = int64(integer32);
	var integer16 int16 = int16(integer32);

	fmt.Println(integer32);
	fmt.Println(integer64);
	fmt.Println(integer16);

	var name = "Leeland";
	var l uint8 = name[0];
	var lString = string(l);

	fmt.Println(name);
	fmt.Println(l);
	fmt.Println(lString);
}