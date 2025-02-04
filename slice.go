package main

import (
	"bytes"
	"fmt"
	"sort"
)

func slice() {

	myslice := []string{"This", "is", "the",
        "tutorial", "of", "Go", "language"}
 
    // Iterate slice
    // using range in for loop
    // without index
    for _, ele := range myslice {
        fmt.Printf("Element = %s\n", ele)
	}

	slice := []int{1, 2, 3}
    slice = append(slice, 4, 5, 6)
 
    fmt.Println("Slice: ", slice) // [1 2 3 4 5 6]

	// Creating an array of size 7
    // and slice this array  till 4
    // and return the reference of the slice
    // Using make function
    var my_slice_1 = make([]int, 4, 7)
    fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
                   my_slice_1, len(my_slice_1), cap(my_slice_1)) // Slice 1 = [0 0 0 0], length = 4, capacity = 7
 
    // Creating another array of size 7
    // and return the reference of the slice
    // Using make function
    var my_slice_2 = make([]int, 7)
    fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
                   my_slice_2, len(my_slice_2), cap(my_slice_2)) // Slice 2 = [0 0 0 0 0 0 0], length = 7, capacity = 7

	// Creating multi-dimensional slice
	s1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s1) // Slice 1 :  [[12 34] [56 47] [29 40] [46 78]]

	// Creating multi-dimensional slice
	s2 := [][]string{
		[]string{"Geeks", "for"},
		[]string{"Geeks", "GFG"},
		[]string{"gfg", "geek"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2) // Slice 2 :  [[Geeks for] [Geeks GFG] [gfg geek]]

	intSlice := []int{42, 23, 16, 15, 8, 4}

	fmt.Println(sort.IntsAreSorted(intSlice))

    // Sort in descending order
    sort.Slice(intSlice, func(i, j int) bool {
        return intSlice[i] > intSlice[j]
    })

    fmt.Println("descending order:", intSlice)

	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
    res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
    res3 := bytes.Trim([]byte("^^Geek&&"), "$")
 
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s", res3)

	data := []byte("a;b;c")

    // Splitting with ';' as the separator
    parts := bytes.Split(data, []byte(";"))
    for _, part := range parts {
        fmt.Println(string(part))  // a b c
    }

}
