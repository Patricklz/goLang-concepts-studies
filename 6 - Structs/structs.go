package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Structs files")

	var u user
	u.name = "Name"
	u.age = 10
	fmt.Println(u)

	adressExample := address{"street name", 22}

	u2 := user{"name2", 12, adressExample}
	fmt.Println(u2)

	u3 := user{name: "name3"}
	fmt.Println(u3)
}
