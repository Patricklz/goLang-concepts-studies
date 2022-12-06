package main

import "fmt"


func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
	
}


func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	sumTotal := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sumTotal)

	write("Hi", 1, 2, 3, 4)
}