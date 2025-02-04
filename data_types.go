package main

import "fmt"

func data_types() {

	var a int16 = 32767
	fmt.Println(a-2, a+2) // 32765 -32767
	var b uint = 32767
	fmt.Println(b-2, b+2) // 32765 32769

	c := 3.14
	d := 13.7
	e := c+d

	fmt.Printf("The type of e is : %T", e) // float64
	fmt.Printf("\nThe value of e is : %f", e) // 16.840000

	var f complex64 = 3 + 5i
	var g complex64 = complex(3 , 5)

	h := f + g

	realNum := real(h)
	imagNum := imag(h)

	fmt.Printf("\nThe value of h is : %v where real part is : %v and imaginary part is : %v \n", h, realNum, imagNum) // (6+10i), 6, 10

	str1 := "Hello, World!"
	str2 := "Hello, Go!"

	res1 := str1 == str2
	res2 := str1 != str2

	fmt.Println(res1, res2) // false true
	fmt.Println(str1 + str2) // Hello, World!Hello, Go!
	fmt.Println(str1[0:5]) // Hello
	fmt.Println(len(str1)) // 13

}