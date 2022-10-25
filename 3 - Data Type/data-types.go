package main

import (
	"errors"
	"fmt"
)

func main() {

	var number int64 = -100000000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	//alias 
	//INT32 = RUNE
	var number3 rune = 12456
	fmt.Println(number3)

	//BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)

	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 12300000000.45
	fmt.Println(realNumber2)

	realNumber3 := 12345.67
	fmt.Println(realNumber3)

	

	var str string = "text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)


	var text string
	fmt.Println(text)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var erro error = errors.New("Internal error")
	fmt.Println(erro)

}
