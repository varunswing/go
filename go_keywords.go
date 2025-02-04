package main

import "fmt";

func keywords() {
	var def1 = "Keywords in Go are reserved words that have special meaning. They cannot be used as identifiers (variable names, function names, etc.";
	var def2 = "The following are the keywords in Go:" ;
	var def3 = "break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto, if, import, interface, map, package, range, " +
			   "return, select, struct, switch, type, var" ;
	var def4 = "Keywords are case-sensitive, meaning that 'break' is a keyword, but 'Break' is not." ;
	var def5 = "Keywords are used to define the syntax and structure of the Go programming language." ;

	fmt.Println(def1);
	fmt.Println(def2);
	fmt.Println(def3);
	fmt.Println(def4);
	fmt.Println(def5);
}