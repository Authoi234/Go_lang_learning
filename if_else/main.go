package main

import (
	"fmt"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	y := 10
	z := 40
	if x > 5 && (y > 5 || z < 43) {
		fmt.Println("How are you?")
	} else {
		fmt.Println("hello")
	}
}