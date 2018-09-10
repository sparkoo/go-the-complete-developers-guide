package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://seznam.cz",
		"http://golang.com",
		"http://stackoverflow.com",
		"http://sparko.cz",
		"http://blabol.czf",
	}
	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is off!")
		return
	}

	fmt.Println(link, "is on")
}
