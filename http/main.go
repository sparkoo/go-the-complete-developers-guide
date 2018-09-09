package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
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

	fmt.Println()

	res, err = http.Get("http://google.com")
	io.Copy(os.Stdout, res.Body)

	fmt.Println()

	res, err = http.Get("http://google.com")
	io.Copy(logWriter{}, res.Body)
}

func (logWriter) Write(p []byte) (n int, err error) {
	fmt.Println("my writer:", string(p))
	fmt.Println("just wrote", len(p), "bytes")
	return len(p), nil
}
