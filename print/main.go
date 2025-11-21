package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.82103

	// Println
	fmt.Println("Age is:", age, "height is:", height, "name is:", name) // give spaces after each print item
	fmt.Println("HELLO WORLD")                                          // print in newline

	// Printf

	fmt.Printf("From Printf Age is: %d", age)
	fmt.Printf("From Printf height is: %.3f\n", height)
	// From Printf Age is: 25From Printf height is: 5.821
	//
	fmt.Printf("Type of Name is: %T\n", name)
	fmt.Printf("Type of height is: %T\n", height)
	fmt.Printf("Type of age is: %T\n", age)

}
