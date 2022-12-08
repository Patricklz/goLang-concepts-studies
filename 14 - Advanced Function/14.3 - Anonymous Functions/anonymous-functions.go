package main


import "fmt"

func main() {

	ret := func(text string) string {
		return fmt.Sprintln("Received -> %s", text)
	}("params")

	fmt.Println(ret)
}