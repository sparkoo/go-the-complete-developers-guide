package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://seznam.cz")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	b := make([]byte, 1024)
	var content []byte

	for i, err := res.Body.Read(b); i > 0 && err == nil; i, err = res.Body.Read(b) {
		if err != nil {
			fmt.Println("Error: ", err)
			res.Body.Close()
			os.Exit(1)
		}

		content = append(content, b[:i]...)
	}
	fmt.Println(string(content))
	res.Body.Close()
}
