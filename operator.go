package main

import (
	"fmt"
	"time"
)

func operator() {
	// 1️⃣ Pointers (& and *)
	x := 42
	ptr := &x  // Get address of x

	fmt.Println("Original value of x:", x) // 42
	fmt.Println("Address of x:", ptr) // 0x1400000e0b0
	fmt.Println("Value at ptr (dereferencing *ptr):", *ptr) // 42

	*ptr = 100 // Change x using pointer
	fmt.Println("New value of x after modifying through pointer:", x) // 100

	// 2️⃣ Channels (<-)
	ch := make(chan int) // Create a channel

	go func() { 
		ch <- *ptr  // Send value of x (100) into channel
	}()

	// Simulate delay to let goroutine execute
	time.Sleep(time.Second)

	receivedValue := <-ch // Receive value from channel
	fmt.Println("Received value from channel:", receivedValue) // 100

	fmt.Println(4 &^ 2) // 4  (&^ clears bits in 4 where 2 has 1s)
}
