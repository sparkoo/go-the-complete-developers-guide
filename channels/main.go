package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		//"http://google.com",
		"http://facebook.com",
		//"http://seznam.cz",
		"http://golang.com",
		//"http://stackoverflow.com",
		//"http://sparko.cz",
		//"http://blabol.czf",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is off!")
		c <- link
		return
	}
	fmt.Println(link, "is on")
	c <- link
}
