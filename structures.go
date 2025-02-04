// Golang program to illustrate
// the pointer to struct
package main

import "fmt"

// Defining a structure
type Employee struct {
	firstName, lastName string
	age, salary		 int
}



// Main Function
func structure() {

	// taking pointer to struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	// emp8.firstName is used to access
	// the field firstName
	fmt.Println("First Name: ", emp8.firstName) // Sam
	fmt.Println("Age: ", emp8.age) // 55

	emp := Employee{
        firstName:    "John Doe",
        age:     30,
        salary: 123,
    }

	fmt.Println(emp.firstName) // John Doe
	emp.firstName = "Varun"
	fmt.Println(emp.firstName) // Varun
}
