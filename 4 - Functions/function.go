package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2

	return sum, subtraction
}

func main() {
	total := sum(10, 20)
	fmt.Println(total)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("text fucntion 1")
	fmt.Println(result)

	// use the underscore to not use the return
	_, substractionResult := calc(10, 15)
	fmt.Println(substractionResult)
}
