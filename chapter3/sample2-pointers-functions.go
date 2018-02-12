package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	changeValue(x)
	fmt.Println(x)

	fmt.Println(x)
	changeValueRef(&x)
	fmt.Println(x)
}

// by value
func changeValue(x int) {
	x = x * x
}

func changeValueRef(x *int) {
	*x = (*x) * (*x)
}
