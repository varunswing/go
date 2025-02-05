package main 

import "fmt"

func myfun(mychnl chan string) { 

	for v := 0; v < 4; v++ { 
		mychnl <- "GeeksforGeeks"
	} 
	close(mychnl) 
} 

func sending(s chan<- string) {
    s <- "GeeksforGeeks"
}

func channel() { 

	// Creating a channel 
	c := make(chan string) 

	go myfun(c) 

	// When the value of ok is set to true means the channel is open and it can send or receive data 
	// When the value of ok is set to false means the channel is closed 
	for { 
		res, ok := <-c 
		if ok == false { 
			fmt.Println("Channel Close ", ok) 
			break
		} 
		fmt.Println("Channel Open ", res, ok) 
	} 

	mychnl := make(chan string) 
  
    // Anonymous goroutine 
    go func() { 
        mychnl <- "GFG"
        mychnl <- "gfg"
        mychnl <- "Geeks"
        mychnl <- "GeeksforGeeks"
        close(mychnl) 
    }() 

    for res := range mychnl { 
        fmt.Println(res) 
    } 

	newChnl := make(chan string, 5) 
    newChnl <- "GFG"
    newChnl <- "gfg"
    newChnl <- "Geeks"
    newChnl <- "GeeksforGeeks"

	fmt.Println("Length of the channel is: ", len(newChnl)) 

	fmt.Println("Capacity of the channel is: ", cap(newChnl)) 

	mychanl := make(chan string)
 
    // Here, sending() function convert the bidirectional channel into send only channel
    go sending(mychanl)
 
    // Here, the channel is sent only inside the goroutine outside the goroutine the channel is bidirectional. So, it print GeeksforGeeks
    fmt.Println(<-mychanl)

} 

// Blocking Send and Receive: In the channel when the data sent to a channel the control is blocked in that send statement until other goroutine reads from that channel. 
// Similarly, when a channel receives data from the goroutine the read statement block until another goroutine statement.
// Zero Value Channel: The zero value of the channel is nil.
// For loop in Channel: A for loop can iterate over the sequential values sent on the channel until it closed.
