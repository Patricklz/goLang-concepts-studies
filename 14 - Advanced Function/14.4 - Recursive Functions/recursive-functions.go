package main

import "fmt"


func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}

func main() {
	fmt.Println("Recursive Function")

	position := uint(15)

	fmt.Println(fibonacci(position))

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}