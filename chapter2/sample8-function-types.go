package main

import "fmt"

func add(x, y int) int { return x + y }
func mul(x, y int) int { return x * y }

func main() {
	var biFunc1 func(int, int) int = add
	biFunc2 := mul

	fmt.Println(biFunc1(1, 2))
	fmt.Println(biFunc2(1, 2))

	fmt.Println(action(10, 20, add))
	fmt.Println(action(10, 20, mul))

	fmt.Println(getAction(1)(10, 20))
	fmt.Println(getAction(2)(10, 20))

	fmt.Println(action(10, 5, func(x, y int) int {
		return x - y
	}))


	f := seqSquare()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func action(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func getAction(t int) func(int, int) int {
	switch t {
	case 1:
		return add
	case 2:
		return mul
	default:
		panic("no such action")
	}
}

func seqSquare() func() int {
	x := 2
	return func() int {
		x++
		return x * x
	}
}
