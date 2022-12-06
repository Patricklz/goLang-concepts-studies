package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array and Slices")

	var array1 [5]string
	array1[0] = "position 1"

	fmt.Println(array1)

	array2 := [5]string{"position 0", "position 1", "position 2", "position 3", "position 4"}
	fmt.Println(array2)

	array3 := [...]int{0, 1, 2, 3, 4}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 17)
	fmt.Println(slice)

	//slice works like a pointer and get the memory reference
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Alter Position"
	fmt.Println(slice2)

	//****Internal Array****
	fmt.Println("------------")
	slice3 := make([]float32, 10, 11) //make creates an array of 15 positions and returns a slice of 10
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacity

	fmt.Println("------------")

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacity
	fmt.Println(slice3)

	fmt.Println("------------")

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacity

	// array is a fixed-length list and alice is a non-fixed-length list

}
