package main

import "fmt"

func main() {

	// What is a map? (HashTable)
	// - key-value pair store
	// - key is unique, value not
	// - length is not limited (dynamic length)

	// Construction: map[keyType]valueType
	var names map[int]string

	// make()-function: allocates and initializes an object of type slice, map, or chan (only).
	names = make(map[int]string)

	// Using a composite literal to initialize a map.
	names = map[int]string{}
	names = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Println("Hello all: ")
	fmt.Println(names)

}
