package main

import (
	"fmt"
	"math"
)

const (
	GFG	 = "GeeksforGeeks"
	Correct = true
	n = 5
	d = 3e10 / n
)

func constants() {

	fmt.Println("value of GFG : ", GFG) // GeeksforGeeks

	fmt.Println("value of Correct : ", Correct) // true

	fmt.Println("value of d : ", (int64(d))) // 6000000000 which is 6e+09

	fmt.Println("value of sin of n : ", math.Sin(n)) // -0.9589242746631385
}
