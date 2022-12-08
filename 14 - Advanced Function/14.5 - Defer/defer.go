package main


import "fmt"

func func1() {
	fmt.Println("exec function 1")
}

func func2() {
	fmt.Println("exec function 2")
}

func studentIsApproved(n1, n2 float32) bool {
	
	defer fmt.Println("Calculated average. The result will be returned")
	fmt.Println("Entering the function to check if the student will be approved")

	avarage := (n1 + n2) / 2

	if avarage >=6 {
		return true
	}

	return false
}

func main() {
	defer func1()

	func2()

	fmt.Println(studentIsApproved(7, 8))
}