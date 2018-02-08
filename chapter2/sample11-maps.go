package main

import "fmt"

func main() {

	people := map[string]int {
		"a": 1,
		"b": 2,
	}

	fmt.Println(people)
	fmt.Println(people["a"])

	if val, ok := people["c"]; ok {
		fmt.Println(val)
	}

	for k, v := range people {
		fmt.Println(fmt.Sprintf("%s %v", k, v))
	}
}