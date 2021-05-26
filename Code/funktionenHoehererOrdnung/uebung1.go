package main

import "fmt"

/**
Übung SW07: Aufgabe 6 Funktionen höherer Ordnung
Implementieren Sie eine Funktion, der man die Liste mit Funktionen und die Seiten eines
Rechteckes übergeben kann, dann die Eigenschaften berechnet und ausgibt.
*/
func main() {

	var x = 10
	var y = 20
	var myListOfFunctions = make([]func(int, int) int, 2, 2)
	myListOfFunctions[0] = area
	myListOfFunctions[1] = circumference

	calcListOfFunctions(myListOfFunctions, x, y)
}

func calcListOfFunctions(list []func(int, int) int, x, y int) {
	for _, function := range list {
		var result = function(x, y)
		fmt.Println("Result of function: ", result)
	}
}

func area(x, y int) int {
	return x * y
}

func circumference(x, y int) int {
	return 2*x + 2*y
}
