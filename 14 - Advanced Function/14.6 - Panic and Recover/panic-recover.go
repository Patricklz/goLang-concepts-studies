package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Recover Execution")
	}
}

func studentIsApproved(n1, n2 float32) bool {
	defer recoverExecution()

	avarage := (n1 + n2) / 2

	if avarage > 6 {
		return true
	} else if avarage < 6 {
		return false
	}

	panic("The average is exactly 6!")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("Post execution")
}
