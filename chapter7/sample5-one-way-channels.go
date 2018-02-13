package main

import "fmt"

func main() {
	// only for sending
	//var intCh chan<- int
	// only for receiving
	//var outCh <-chan int
	intCh := make(chan int, 2)
	go f(5, intCh)
	fmt.Println(<-intCh)
}

func f(n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	// compile error - send only channel
	//<- ch
	ch <- result
}
