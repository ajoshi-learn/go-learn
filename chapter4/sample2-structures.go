package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var tom person
	var tom2 person = person{"Tom", 27}
	bob := person{"Bob", 22}

	fmt.Println(tom)
	fmt.Println(tom2)
	fmt.Println(bob)
}
