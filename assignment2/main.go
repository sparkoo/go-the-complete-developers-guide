package main

import "fmt"

type triangle struct{}
type square struct{}

type shape interface {
	getArea() float64
}

func main() {
	printArea(triangle{})
	printArea(square{})
}

func (t triangle) getArea() float64 {
	return 25.7
}

func (s square) getArea() float64 {
	return 12.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
