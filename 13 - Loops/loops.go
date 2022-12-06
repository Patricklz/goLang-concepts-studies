package main

import (
	"fmt"
	// "time"
)


func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Increment i")
	// 	time.Sleep(time.Second)
		
	// }


	// for j := 0; j< 10; j++ {
	// 	fmt.Println("Increment j", j)
	// 	time.Sleep(time.Second)
	// }

	names := [3]string{"Jonh", "Tony", "Bruce"}

	for index, name := range names {
		fmt.Println(index, name)

	}


}