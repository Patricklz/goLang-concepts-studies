package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(1)
	generic(true)

	fmt.Println(1, 2, "String", false, true, float64(12345))

	map1 := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(map1)
}
