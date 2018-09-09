package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
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

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%v => %v\n", k, v)
	}
}
