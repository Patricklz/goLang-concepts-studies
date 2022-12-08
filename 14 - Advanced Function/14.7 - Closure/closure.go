package main

import "fmt"


func closureFunction() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	mainText := "Inside main"
	fmt.Println(mainText)

	newFunction := closureFunction()

	newFunction()
}
