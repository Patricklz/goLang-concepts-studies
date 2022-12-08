package main

import "fmt"

var n int

// this function exec first before main function
func init() {
	fmt.Println("Exec init function")
	n = 10
}

func main() {
	fmt.Println("Exec main function")
	fmt.Println(n)
}
