package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)
 
func str() {

    myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
 
    mystring1 := string(myslice1)
 
    fmt.Println("String 1: ", mystring1) // Geeks
 
    // Creating and initializing a slice of rune
    myslice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0}

	mystring2 := string(myslice2)

    for index, s := range mystring2 {
        fmt.Printf("The index number of %c is %d\n", s, index) 
    }

	fmt.Println(utf8.RuneCountInString(mystring2)) // 4 (length of string)

	s := "@@Hello, Geeks!!"
    result := strings.Trim(s, "@!") // Removes all specified leading and trailing characters from a string.
    fmt.Println(result) // Hello, Geeks

	s = "@@Hello, Geeks!!"
    result = strings.TrimLeft(s, "@!") // Removes all specified leading characters from a string.
    fmt.Println(result) // Hello, Geeks!!

	s = "@@Hello, Geeks"
    result = strings.TrimRight(s, "Geeks") // Removes all specified trailing characters from a string.
    fmt.Println(result) // @@Hello,
 
	s = "   Hello, Geeks   "
    result = strings.TrimSpace(s) // Removes all leading and trailing white spaces from a string.
    fmt.Println(result) // Hello, Geeks

	s = "@@Hello, Geeks!!"
    result = strings.TrimPrefix(s, "@@") // Removes the specified prefix from a string.
    fmt.Println(result) // Hello, Geeks!!

	s = "@@Hello, Geeks"
    result = strings.TrimSuffix(s, "eeks") // Removes the specified suffix from a string.
    fmt.Println(result) // @@Hello, G

	s = "Welcome,to,GeeksforGeeks"

    resSlice := strings.Split(s, ",") // Splits a string into substrings based on the separator.
    fmt.Println("Result:", resSlice) // [Welcome to GeeksforGeeks]

	resSlice = strings.SplitAfter(s, ",") // Splits a string into substrings after each instance of the separator.
    fmt.Println("Result:", resSlice) // [Welcome, to, GeeksforGeeks]

	resSlice = strings.SplitN(s, ",", 2) // Splits a string into substrings based on the separator and returns a slice of the first n substrings.
    fmt.Println("Result:", resSlice) // [Welcome to, GeeksforGeeks]

	resSlice = strings.SplitAfterN(s, ",", 2) // Splits a string into substrings after each instance of the separator and returns a slice of the first n substrings.
	fmt.Println("Result:", resSlice) // [Welcome, to, GeeksforGeeks]

	s1 := "Hello"
    s2 := "Geeks"
    s3 := "Hello"

	fmt.Println("s1 == s2:", s1 == s2) // false
    fmt.Println("s1 == s3:", s1 == s3) // true
    fmt.Println("s1 != s2:", s1 != s2) // true
	fmt.Println("s1 < s2:", s1 < s2)   // true
    fmt.Println("s1 > s2:", s1 > s2)   // false
    fmt.Println("s1 <= s3:", s1 <= s3) // true
    fmt.Println("s1 >= s3:", s1 >= s3) // true

	fmt.Println("strings.Compare(s1, s2):", strings.Compare(s1, s2)) // -1
    fmt.Println("strings.Compare(s1, s3):", strings.Compare(s1, s3)) // 0
    fmt.Println("strings.Compare(s2, s1):", strings.Compare(s2, s1)) // 1
}
