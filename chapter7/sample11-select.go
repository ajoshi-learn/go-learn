package main

import "fmt"

func main() {

	first := make(chan int)
	second := make(chan int)

	go func() {
		first <- 1
	}()

	go func() {
		second <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-first:
			fmt.Println("first received ", msg)
		case msg := <-second:
			fmt.Println("second received", msg)
		}
	}
}
