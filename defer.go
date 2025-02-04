package main

import "fmt"

// Functions
func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}

// Main function
func defer_keyword() {

	fmt.Println("Start")
	defer fmt.Println("End")
	defer add(34, 56)
	defer add(10, 10) 

	// Start
	// Result:  20
	// Result:  90
	// End
}

// In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order.
// In the defer statements, the arguments are evaluated when the defer statement is executed, not when it is called.
// Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.
