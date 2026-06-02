package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	fmt.Printf("Numbers has data type : %T\n", numbers)
	fmt.Println("Length of numbers:", len(numbers)) 
	numbers = append(numbers, 63,134,6,4,5,56,7,3,56,5,10,3,2)
	fmt.Println("Length of numbers:", len(numbers)) 

	other_numbers  := make([]int, 3, 5)
	fmt.Println("slice:", other_numbers)
	fmt.Println("length", len(other_numbers))
	fmt.Println("capacity", cap(other_numbers))

	other_numbers = append(other_numbers, 1,2)
	fmt.Println("slice:", other_numbers)
	fmt.Println("length", len(other_numbers))
	fmt.Println("capacity", cap(other_numbers))

	other_numbers = append(other_numbers, 6)
	fmt.Println("slice:", other_numbers)
	fmt.Println("length", len(other_numbers))
	fmt.Println("capacity", cap(other_numbers))


}