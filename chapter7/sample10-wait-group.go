package main

import (
	"sync"
	"fmt"
)

func main() {
	// countdownlatch
	var wg sync.WaitGroup
	wg.Add(2)
	counter := 5
	doubleCounter := func() {
		defer wg.Done()
		counter *= 2
	}

	go doubleCounter()
	go doubleCounter()

	wg.Wait()
	fmt.Println(counter)
}