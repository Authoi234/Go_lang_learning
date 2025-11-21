package main

import "fmt"

func simpleFunc() {
	fmt.Println("This is a simple function")
}

func add(a, b int) (result int) {
	// return a+b
	result = a + b
	return
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	simpleFunc()
	ans := add(3, 4)
	ans2 := multiply(3, 4)
	fmt.Println("Addition is:", ans)
	fmt.Println("Multiplication is:", ans2) 
}
