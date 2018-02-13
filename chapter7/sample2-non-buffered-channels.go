package main

import "fmt"

func main() {

	intCh := make(chan int)
	go fact(5, intCh)

	fmt.Println(<-intCh)
}

func fact(n int, ch chan int) {
	if (n < 1) {
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
