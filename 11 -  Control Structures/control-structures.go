package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := -5

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less or equal than 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Number is greater than zero")
	} else {
		fmt.Println("Number is less than zero")
	}
}
