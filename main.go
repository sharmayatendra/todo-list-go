package main

import (
	"fmt"
	"net/http"
)

func main() {
	// using variables
	var firstVariable = "first variable"

	// alternative declaration syntax for variables
	secondVariable := "second variable"

	thirdVariable := "third variable"

	// slices syntax:
	combinedVariables := []string {firstVariable, secondVariable, thirdVariable }
	
	// so diff b/w array & slice is array is fixed size while slice is dynamically allocated
	// example:
	// listArray := [2]string {firstVariable, secondVariable, secondVariable}
	// above code will throw an error as it is strict for 2 elements
	// fmt.Println(combinedVariables)

	printLoopUsingFunction(combinedVariables)
	combinedVariables = addNewTaskToList(combinedVariables, "New task")
	combinedVariables = addNewTaskToList(combinedVariables, "first go code")

	fmt.Println()
	fmt.Println("After adding new task")
	fmt.Println()

	// loops syntax:
	// for index, task := range combinedVariables {
	// 	fmt.Printf("%d. %s\n", index + 1, task)
	// }
	printLoopUsingFunction(combinedVariables)

	http.HandleFunc("/hello", helloUserHandler);

	http.ListenAndServe(":8080", nil)
}

func helloUserHandler(w http.ResponseWriter, r *http.Request) {
	greetUser := "Hello, User!"
	fmt.Fprintln(w, greetUser)
}

// function
func printLoopUsingFunction(params []string) {
	for index, task := range params {
		fmt.Printf("%d. %s\n", index + 1, task)
	}
}

// adding new task to list using append function
// what func takes parameters & return value
func addNewTaskToList(currentTaskList []string, newTask string) []string {
	updatedTaskList := append(currentTaskList, newTask)

	return updatedTaskList
}