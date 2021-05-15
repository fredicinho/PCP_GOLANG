package main

import "fmt"

func main() {

	var x int = 10
	var y int = 20
	var myListOfFunctions = make([]func(int, int) int, 2, 2)
	myListOfFunctions[0] = area
	myListOfFunctions[1] = circumference

	var results = []int{}
	for _, function := range myListOfFunctions {
		results = append(results, function(x, y))
	}

	fmt.Println(results)
}

func area(x, y int) int {
	return x * y
}

func circumference(x, y int) int {
	return 2*x + 2*y
}
