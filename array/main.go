package main

import "fmt"

func main(){
	fmt.Println("We are learning arrays")

	var names [5]string

	// 0 1 2 3 4
	names[0] = "Abul"
	names[1] = "Habul"

	fmt.Println("Names:", names)
	
	var numbers = [6]int{1,2,3,4,5} // 1,2,3,4,5,0
	fmt.Println("Numbers:", numbers)
	fmt.Println("length:", len(numbers)) // 6
	
	fmt.Println("value of name at 2nd index:", names[1])
	fmt.Printf("Names: %q \n", names)
}