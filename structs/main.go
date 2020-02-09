package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	myPerson := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "a@b.com",
			zipCode: 94000,
		},
	}
	myPerson.print()
	myPerson.updateName("Frank")
	myPerson.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
