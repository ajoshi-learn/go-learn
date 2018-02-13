package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 1
	intCh <- 2
	intCh <- 3

	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
}
