package main

import "fmt"

func main() {

	printSum(1, 2)

}

func printSum(x, y int) {
	fmt.Println(x + y)
}

func addVararg(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func addNamed(numbers ...int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}

func severalValues(numbers ...int) (mul, sum int) {
	mul = 1
	for _, v := range numbers {
		sum += v
		mul *= v
	}
	return
}