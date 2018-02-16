package main

import (
	"os"
	"fmt"
)

const FILENAME = "/tmp/test.txt"

func main() {
	file, err := os.Create(FILENAME)
	if err == nil {
		fmt.Println(file.Name())
	}
	// os.OpenFile() - opens or creates if file not exists
	// f1, err := os.OpenFile("sometext.txt", os.O_RDONLY, 0666)
	opened, err := os.Open(FILENAME)
	opened.Close()
}
