package main

import "fmt"

func main() {

	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 10 * 5
	restDivision := 10 % 2

	fmt.Println(sum, subtraction, division, multiplication, restDivision)

	//can't compare var with different types
	var number1 int16 = 10
	var number2 int16 = 25
	result := number1 + number2
	fmt.Println(result)

	//attribution
	var variable1 string = "String"
	variable2 := "String2"
	fmt.Println(variable1, variable2)

	//relational operators
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// logical operators
	fmt.Println("--------------------")
	true1, false1 := true, false

	fmt.Println(true1 && false1)
	fmt.Println(true1 || false1)
	fmt.Println(!true1)
	fmt.Println(!false1)

	// unary operators
	number := 10
	number++
	number += 15 //number = number + 15
	fmt.Println(number)

	number--
	number -= 20 // number = number - 20

	number *= 3  // number = number * 3
	number /= 10 // number = number / 10
	number %= 3

	fmt.Println(number)

	var text string

	if number > 5 {
		text = "Greater than 5"
	} else {
		text = "Less than 5"
	}

	fmt.Println(text)

}
