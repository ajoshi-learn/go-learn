package main

import "fmt"

func main() {
	fmt.Println(divide(6, 3))
	fmt.Println(divide(6, 0))
}

func divide(x, y int) int {
	if y == 0 {
		panic("Division by zero")
	}
	return x / y
}
