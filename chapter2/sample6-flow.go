package main

import "fmt"

func main() {
	a := 6
	b := 5

	if a < b {
		fmt.Println("a less than b")
	}

	if a < b {
		fmt.Println("a less than b")
	} else {
		fmt.Println("b less than a")
	}

	switch(a) {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3,4,5:
		fmt.Println("other")
	default:
		fmt.Println("default")
	}

	var i = 1
	for ; i < 10; i++{
		fmt.Println(i * i)
	}

	var i1 = 1
	for ; i1 < 10;{
		fmt.Println(i1 * i)
		i1++
	}

	var i2 = 1
	for i2 < 10{
		fmt.Println(i2 * i)
		i2++
	}
}
