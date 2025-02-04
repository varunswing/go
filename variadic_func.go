package main
import "fmt"

// Variadic function to calculate sum
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Function with a regular parameter and a variadic parameter
func greet(prefix string, nums ...int) {
    fmt.Println(prefix)
    sum := sum(nums...)
	fmt.Println(sum)
}
func variadic_func() {
    greet("Sum of numbers:", 1, 2, 3)
    greet("Another sum:", 4, 5)
    greet("No numbers sum:")
}

// Variadic functions can only have one variadic parameter, and it must be the last parameter.
// You cannot have multiple variadic parameters in a single function definition.