package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {

}

func (c *Car) move() {
	fmt.Println("Moving")
}

func main() {
	var veh Vehicle = new(Car)
	veh.move()
}
