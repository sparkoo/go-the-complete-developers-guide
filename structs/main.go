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
	fmt.Println(john)

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   contactInfo{email: "alex@anderson", zip: 123456}}
	fmt.Println(alex)

	var jane person
	fmt.Println(jane)
	fmt.Printf("%+v\n", jane)
	jane.firstName = "Jane"
	jane.lastName = "Doe"
	jane.contact.email = "jane@doe"
	jane.contact.zip = 123456
	fmt.Printf("%+v\n", jane)
}
