package main

import "fmt"

type mile int
type km int

type library []string

func main() {
	var distance mile = 5
	printDistanceMiles(distance)

	//var distanceKm km = 6
	// compile error - type check
	//printDistanceMiles(distanceKm)

	myLib := library{"B1", "B2"}
	printBooks(myLib)
}

func printDistanceMiles(distance mile) {
	fmt.Println(fmt.Sprintf("%d miles", distance))
}

func printBooks(books library) {
	for _, v := range books {
		fmt.Println(v)
	}
}
