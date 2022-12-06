package main

import "fmt"


func weekDays(number int) string {
	switch number {
	case 1: 
		return "sunday"
	case 2: 
		return "monday"
	case 3: 
		return "thursday"
	default:
		return "invalid number"
	}
}

func weekDays2(number int) string {
	var day string
	
	switch number {
	case 1: 
		day = "sunday"
	case 2: 
		day = "monday"
	case 3: 
		day = "thursday"
	default:
		day = "invalid number"
	}
}
	

func main() {
	fmt.Println("Switch")

	day := weekDays(1)
	fmt.Println(day)

}