package main

import "fmt"

func main() {
	// using variables
	var firstVariable = "first variable"

	// alternative declaration syntax for variables
	secondVariable := "second variable"

	// slices syntax:
	combinedVariables := []string {firstVariable, secondVariable }
	
	// so diff b/w array & slice is array is fixed size while slice is dynamically allocated
	// example:
	// listArray := [2]string {firstVariable, secondVariable, secondVariable}
	// above code will throw an error as it is strict for 2 elements
	fmt.Println(combinedVariables)
}
