package main

import "fmt"

type library1 []string

type person1 struct {
	name string
	age  int
}

func main() {
	var lib library1 = library1{"Book1", "Book2"}
	lib.print()
}

func (l library1) print() {
	for _, v := range l {
		fmt.Println(v)
	}
}

func (p person1) print() {
	fmt.Println(p.age)
	fmt.Println(p.name)
}

func (p *person1) incrementAge() {
	p.age++
}
