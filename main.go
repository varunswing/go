// Declaration of the main package
package main

// Importing packages
import (
    "fmt"
    "sort"
    "strings"
    "time"
)

func main() {
	// print()
	// keywords()
	// data_types()
	// variables()
	// constants()
	// operator()
	// if_else()
	// for_loop()
	// switch_case()
	// functions()
	// variadic_func()
	// anonymous_func()
	// defer_keyword()
	// methods()
	// structure()
	// nested_structures()
	// structures_anonymous()
	// array()
	// slice()
	// str()
	// pointer()
	// go_routine()
	// select_go()
	// channel()
	interfaceExample()

    s := []int{345, 78, 123, 10, 76, 2, 567, 5}
    sort.Ints(s)
    fmt.Println("Sorted slice: ", s)  // Sorting the given slice

    
    fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))  // Finding the index

    
    fmt.Println("Time: ", time.Now().Unix()) // Finding the time

	// Only for receiving
    mychanl1 := make(<-chan string)
 
    // Only for sending
    mychanl2 := make(chan<- string)
 
    // Display the types of channels
    fmt.Printf("%T", mychanl1)
    fmt.Printf("\n%T", mychanl2)
}