package main

import "fmt"

func variables(){

	var1, var2, var3 := 1, "bac", 3.14
	var var4 int
	var var5 string
	var var6 float64
	var var7 bool
	var var8 complex64
	var var9, var10, var11 = 1, "bac", 3.14;

	fmt.Println(var1, var2, var3) // 1 bac 3.14
	fmt.Println(var4, var5, var6, var7, var8) // 0 "" 0 false (0+0i)
	fmt.Println(var9, var10, var11) // 1 bac 3.14
}