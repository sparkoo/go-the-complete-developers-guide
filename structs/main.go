package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	john := person{"John", "Doe"}
	fmt.Println(john)

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var jane person
	fmt.Println(jane)
	fmt.Printf("%+v\n", jane)
	jane.firstName = "Jane"
	jane.lastName = "Doe"
	fmt.Printf("%+v\n", jane)
}
