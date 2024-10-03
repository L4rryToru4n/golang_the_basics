package main

import "fmt"

type Person struct {
	Name       string
	JoinedAt   string
	MemberType string
}

func (person Person) GetInfo() [3]string {

	dataArr := [3]string{
		person.Name,
		person.JoinedAt,
		person.MemberType,
	}

	return dataArr
}

func (person Person) GetName() string {
	return person.Name
}

type PrintCustomer interface {
	GetInfo() [3]string
}

type PrintCustomerName interface {
	GetName() string
}

func getCustomerInfo(customer PrintCustomer) {
	fmt.Println("DATA:", customer.GetInfo())
}

func getCustomerName(customer PrintCustomerName) {
	fmt.Println("NAME:", customer.GetName())
}

func main() {

	person := Person{Name: "Kara", JoinedAt: "03-23-2007", MemberType: "Regular"}
	
	getCustomerInfo(person)
	getCustomerName(person)
}
