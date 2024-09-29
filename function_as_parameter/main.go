package main

import "fmt"

func getPersonel(name string) string {
	if name != "Kara 'Starbuck' Thrace" {
		return "Good bye " + name
	} else {
		return "Farewell " + name
	}
}

func getGoodBye(name string, personel func (string) string) {
	personelName := personel(name)
	
	fmt.Println(personelName)
}

func main() {
	personelKara := "Kara 'Starbuck' Thrace"
	personelLee := "Lee 'Apollo' Adama"

	getGoodBye(personelKara, getPersonel)

	personel := getPersonel

	getGoodBye(personelLee, personel)
}