package main

import (
    "fmt"
    "myLearning/myutil"
)

func main() {
	fmt.Println("Hello, World!")
	myutil.PrintMessage("I am Authoi")

	// variable with type and notype
	var name string = "Authoi in variable"
	var version float64 = 1.0
	fmt.Println(name)
	fmt.Println(version)
	var mynum = 7 // unspecific type variable
	var decided bool = false
	fmt.Println(mynum)
	fmt.Println(decided)

	// constant
	const pi = 3.14159265359
	// pi = 4.0 // constant cannot be changed
	fmt.Println(pi)

	//  Short hand Variable Declaration

	person := "Md Jawad Jabbar Khan Authoi"
	age := 13
	fmt.Println(person, ".age is:", age)

	// Public vs private that can be accessed or unaccessed outside the package
	var Public = "I am accessable from outside the package"
	var private = "I am unaccessable from outside the package"
	fmt.Println(Public)
	fmt.Println(private)
}
