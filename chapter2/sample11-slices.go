package main

import "fmt"

func main() {
	//var users []string

	numbers := []int{1, 2, 3}

	for _, v := range numbers {
		fmt.Println(v)
	}

	newNum := make([]int, 8)
	for _, v := range newNum {
		fmt.Println(v)
	}
}

func addNewElement(slice []int, element int) {
	slice = append(slice, element)
}

func removeElement(slice []int, index int) {
	slice = append(slice[:index], slice[index+1:]...)
}
