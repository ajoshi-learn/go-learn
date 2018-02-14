package main

import (
	"sync"
	"fmt"
)

var counter = 0

func main() {
	var mutex sync.Mutex
	bChan := make(chan bool)
	for i := 0; i < 5; i++ {
		go work(i, bChan, &mutex)
	}
	for i := 0; i < 5; i++ {
		<-bChan
	}
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for i := 1; i <= number; i++ {
		counter++
		fmt.Println(number, " - ", counter)
	}
	mutex.Unlock()
	ch <- true
}
