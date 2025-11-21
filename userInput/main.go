package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, Whats your name?")

	// Scan function
	// var name string;
	// fmt.Scan(&name)  // Read input till space
	// fmt.Println("Hello sir,",name)

	// Bufio Reader

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello sir,", name)

}
