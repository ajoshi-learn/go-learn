package main

import "fmt"

func main() {

	var x int = 4
	var p *int
	var empty *int

	p = &x
	// address of x
	fmt.Println(p)
	// value of x
	fmt.Println(*p)
	// value of p
	fmt.Println(&p)

	fmt.Println(empty)

	pointer := new(int)
	// address of x
	fmt.Println(pointer)
	// value of x
	fmt.Println(*pointer)
}
