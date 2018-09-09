package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	colors3["white"] = "#ffffff"
	fmt.Println(colors3)

	delete(colors3, "white")
	fmt.Println(colors3)
}
