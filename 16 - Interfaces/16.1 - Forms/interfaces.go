package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

// this function receive a form and this must have a method with the same of interface
func writeArea(f form) {
	fmt.Printf("The area of the shape is %0.2f\n", f.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
