package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	close(intCh)
	// panic - channel is closed
	// intCh <- 10
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)

	intCh1 := make(chan int, 3)
	intCh1 <- 10
	intCh1 <- 3
	close(intCh1)

	for i := 0; i < cap(intCh1); i++ {
		if val, opened := <-intCh1; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}

	}
}
