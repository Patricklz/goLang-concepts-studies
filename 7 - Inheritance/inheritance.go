package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
}

type student struct {
	person
	university string
	course     string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"name", "last name", 10}
	fmt.Println(p1)

	s1 := student{p1, "Law", "MIT"}
	fmt.Println(s1)
	fmt.Println(s1.name)
}
