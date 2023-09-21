package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	// we need to close, using defer statement it'll execute after all code
	defer req.Body.Close()
}
