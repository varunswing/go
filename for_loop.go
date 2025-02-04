package main 

import "fmt"

func for_loop() { 
	
	// for loop with single condition
	for i := 0; i < 2; i++{ 
		fmt.Printf("Varun\n") // Varun Varun
	} 

	// for loop with infinite condition
	for {
		fmt.Println("Hello") // Hello
		break
	}

	// for loop as while loop
	i:= 0 
    for i < 3 { 
       i += 2 
    }
	fmt.Println(i) // 4

	rvariable:= []string{"GFG", "Geeks", "GeeksforGeeks"}  
      
    // i and j stores the value of rvariable 
    // i store index number of individual string and 
    // j store individual string of the given array 
    for i, j:= range rvariable { 
       fmt.Println(i, j)  // 0 GFG 1 Geeks 2 GeeksforGeeks
    }

	// for loop with range clause
	for i, j:= range "XabCd" { 
		fmt.Printf("The index number of %U is %d\n", j, i)   // The index number of U+0058 is 0 The index number of U+0061 is 1 The index number of U+0062 is 2 The index number of U+0043 is 3 The index number of U+0064 is 4
	}

	// for loop with map
	mmap := map[int]string{ 
        22:"Geeks", 
        33:"GFG", 
        44:"GeeksforGeeks", 
    } 
    for key, value:= range mmap { 
       fmt.Println(key, value)  // 22 Geeks 33 GFG 44 GeeksforGeeks
    }

	// for loop with channel
	chnl := make(chan int) 
    go func(){ 
        chnl <- 100 
        chnl <- 1000 
       chnl <- 10000 
       chnl <- 100000 
       close(chnl) 
    }() 
    for i:= range chnl { 
       fmt.Println(i)  // 100 1000 10000 100000
    }

	
} 
