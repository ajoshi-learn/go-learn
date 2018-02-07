package main

import "fmt"

func main() {
	const pi float64 = 3.14

	const (
		pi1 float64 = 3.1415
		e float64 = 2.7182
	)

	var hello string
	hello = "1"

	var a, b, c string

	// explicit
	var explicit = "explicit"

	// implicit
	var implicit = "implicit"

	// short
	short := "short"

	// multiple assignment
	var (
		name = "name"
		age  = 21
	)

	var name1, age1 = "name1", 21

	fmt.Println(hello)
	fmt.Println(a, b, c)
	fmt.Println(explicit)
	fmt.Println(implicit)

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(name1)
	fmt.Println(age1)

	fmt.Println(short)
}
