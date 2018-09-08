package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	john := person{"John", "Doe", contactInfo{"john@doe", 123456}}
	john.print()

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   contactInfo{email: "alex@anderson", zip: 123456}}
	alex.print()

	var jane person
	jane.print()
	jane.firstName = "Jane"
	jane.lastName = "Doe"
	jane.contact.email = "jane@doe"
	jane.contact.zip = 123456
	jane.print()

	jane.updateName("Jasmine")
	jane.print()

	jane.updateNamePointer("Jasmine")
	jane.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNamePointer(newFirstName string) {
	p.firstName = newFirstName
}
