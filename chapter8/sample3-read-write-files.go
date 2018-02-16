package main

import (
	"os"
	"fmt"
	"io"
)

const FILENAME1 = "/tmp/test"

func main() {
	text := "Simple text"
	file, err := os.Create(FILENAME1)

	if err != nil {
		fmt.Println("Unable to open file", err)
		os.Exit(1)
	}
	file.WriteString(text)
	fmt.Println("Done")
	file.Close()

	file, err = os.Open(FILENAME1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	for {
		n, readError := file.Read(data)
		if readError == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}
