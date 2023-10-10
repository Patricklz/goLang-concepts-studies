package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World!") //goroutine
	write("Coding in Go!")

}

func write(text string) {
	for i := 0; i < 5; i++{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
