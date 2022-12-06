package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	var var1 int = 10
	var var2 int = var1

	var1++

	fmt.Println(var1, var2)

	var var3 int
	var pointer *int

	var3 = 100
	pointer = &var3

	fmt.Println(var3, *pointer)

}
