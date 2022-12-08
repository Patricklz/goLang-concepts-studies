package main

import (
	"fmt"
)

func inverterSignal(number int) int {
	return number * -1
}

func inventerSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	inverterNumber := inverterSignal(number)
	fmt.Println(inverterNumber)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	inventerSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
