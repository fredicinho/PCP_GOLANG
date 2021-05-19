package main

import "fmt"

func main() {

	// Array
	// - is not resizable
	// - value type
	//

	// What are slices?
	// - resizable arrays (darunter ist ein Array)
	// - reference-type

	// Declaration
	var mySlice1 []string

	// Initialization
	// with literal Syntax
	var mySlice2 = []string{"This", "is", "a", "string", "slice"}
	// with make() --> creates empty slice of length 5
	var mySlice3 = make([]string, 5)

	// Create Slice from array
	var myArray = [4]int{1, 2, 3, 4}
	var mySliceFromArray1 = myArray[1:3] // --> Gets elements of index 1 until 3 (without 3)
	var mySliceFromArray2 = myArray[0:]  // --> Gets all elements from Array

	// Append element into slice
	mySliceFromArray2 = append(mySliceFromArray2, 3)

	// Access and modify slice --> access with index like an array
	mySliceFromArray2[0] = 10

	// Length and Capacity of Slices
	// Length: How many items are inside the slice
	// Capacity: The capacity of the underlying Array
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	s = s[:5]     // panic: runtime error:
	printSlice(s) // slice bounds out of range [:5] with capacity 4

	// Slice is a ReferenceType!
	// If passed into a function, changes on it will also be recognized outside that function!

	// Multidimensional Slices
	var myMultiDimensionalSlice = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 8},
	}

	// Iterating over a slice
	var myIteratingSlice = []string{"One", "Two", "Three"}
	for index, value := range myIteratingSlice {
		fmt.Println(index, value)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
