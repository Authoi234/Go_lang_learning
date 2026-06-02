package main

import (
	"fmt"
)

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Unknown day")
	}

	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}
}