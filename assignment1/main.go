package main

import "fmt"

func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	for no := range numbers {
		if no % 2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	}
}
