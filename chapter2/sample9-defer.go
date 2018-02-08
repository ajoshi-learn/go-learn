package main

import "fmt"

func main() {
	defer finish1()
	defer finish2()

	fmt.Println("Start")
	fmt.Println("In progress")
}

func finish1() {
	fmt.Println("Finish1")
}

func finish2() {
	fmt.Println("Finish2")
}