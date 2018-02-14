package main

import "fmt"

func main() {
	results := make(map[int]int)
	syncChannel := make(chan struct{})
	go fsync(10, syncChannel, results)

	<-syncChannel
	for k, v := range results {
		fmt.Println(fmt.Sprintf("results[%d] = %d", k, v))
	}
}

func fsync(n int, ch chan struct{}, results map[int]int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
