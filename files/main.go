package main

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("hello world"))

	validateError(err)

	fmt.Printf("Created file with %d bytes\n", size)

	f.Close()

	// read
	file, err := os.ReadFile("file.txt")

	validateError(err)

	fmt.Println(string(file))

	// read by pieces
	file2, err := os.Open("file.txt")
	validateError(err)

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 20)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	validateError(err)
}

func validateError(err error) {
	if err != nil {
		panic(err)
	}
}
