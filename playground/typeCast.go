package main

type color string

func main() {
	c := color("Red")
	println(c.describe("is an awesome color"))
}

func (c color) describe(description string) (string) {
	return string(c) + " " + description
}
