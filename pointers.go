package main

import "fmt"

func ptf(a *int) { 
  
    *a = 748
} 

type Individual struct {
    name string
    age  int
}

func pointer() {

	var y = 458
	
	var p = &y

	fmt.Println("Value stored in y before changing = ", y) // 458
	fmt.Println("Address of y = ", &y) // 0xc0000b6010
	fmt.Println("Value stored in pointer variable p = ", p) // 0xc0000b6010

	fmt.Println("Value stored in y(*p) Before Changing = ", *p) // 458

	*p = 500 // Changing the value stored in y

	fmt.Println("Value stored in y(*p) after Changing = ", y) // 500

	ptf(&y)

	fmt.Println(y) // 748

    personPointer := new(Individual)
    personPointer.name = "B"
    personPointer.age = 30

	fmt.Println("Struct created with new:", *personPointer) // {B 30}

	fmt.Println("Name:", personPointer.name) // Automatically dereferences
    fmt.Println("Age:", personPointer.age)

    // Modifying struct values using the pointer
    personPointer.age = 26

	p1 := Individual{name: "A", age: 25}
    fmt.Println("Updated struct:", p1) // {A 25}

	var V int = 100 
       
    var pt1 *int = &V 
       
    var pt2 **int = &pt1 
   
    fmt.Println("The Value of Variable V is = ", V) // 100
    fmt.Println("Address of variable V is = ", &V) // 0xc0000b6018
   
    fmt.Println("The Value of pt1 is = ", pt1) // 0xc0000b6018
    fmt.Println("Address of pt1 is = ", &pt1) // 0xc0000b6020
   
    fmt.Println("The value of pt2 is = ", pt2) // 0xc0000b6020
   
    fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2) // 0xc0000b6018
       
    fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2) // 100

	*pt1 = 200 
  
    fmt.Println("Value stored in v after changing pt1 = ", V) // 200
  
    // changing the value of v by assigning 
    // the new value to the pointer pt2 
    **pt2 = 300 
  
    fmt.Println("Value stored in v after changing pt2 = ", V) // 300

}
