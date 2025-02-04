package main
import "fmt"

func multiplyByValue(a, b int) int {
    a = a * 2 // modifying a inside the function
    return a * b
}

func multiplyByReference(a, b *int) int {
    *a = *a * 2 // modifying a's value at its memory address
    return *a * *b
}

func functions() {

    x := 5
    y := 10
    fmt.Printf("Before: x = %d, y = %d\n", x, y)
    result := multiplyByValue(x, y)
    fmt.Printf("multiplication: %d\n", result)
    fmt.Printf("After: x = %d, y = %d\n", x, y) // x = 5, y = 10 (value not changed)

	x = 5
    y = 10
    fmt.Printf("Before: x = %d, y = %d\n", x, y)
    result = multiplyByReference(&x, &y)
    fmt.Printf("multiplication: %d\n", result)
    fmt.Printf("After: x = %d, y = %d\n", x, y) // x = 10, y = 10 (value changed)
}