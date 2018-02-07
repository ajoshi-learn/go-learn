package main

import "fmt"

func main() {
	var numbers [5]int
	var numbersExplicit [5]int = [5]int{1, 2, 3, 4, 5}
	var numbersImplicit = [5]int{1, 2, 3, 4, 5}
	numbersImplicit2 := [5]int{1, 2, 3, 4, 5}
	numbersImplicit3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	fmt.Println(numbersExplicit)
	fmt.Println(numbersImplicit)
	fmt.Println(numbersImplicit2)
	fmt.Println(numbersImplicit3)

	// ! Array length is a part of type !


	fmt.Println(numbersImplicit[0])

	for i, v := range numbersImplicit {
		fmt.Println(i, v)
	}
}
