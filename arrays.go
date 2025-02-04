package main 

import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func calculateAverage(arr [5]int, size int) int {
    var sum int
    for _, value := range arr {
        sum += value
    }
    return sum / size
}

func array() { 

	arr1:= [3]int{9,7,6} 
	arr2:= [...]int{9,7,6} 
	arr3:= [3]int{9,5,3} 

	fmt.Println(arr1==arr2) 
	fmt.Println(arr2==arr3) 
	fmt.Println(arr1==arr3) 

	// This will give and error because the 
	// type of arr1 and arr4 is a mismatch 
	/* 
	arr4:= [4]int{9,7,6} 
	fmt.Println(arr1==arr4) 
	*/

	var destination *[5]int = &source

    fmt.Println("Source:", source)
    fmt.Println("Destination Array via pointer:", *destination)

	average := calculateAverage(source, len(source))
    fmt.Printf("%d\n", average)
} 
