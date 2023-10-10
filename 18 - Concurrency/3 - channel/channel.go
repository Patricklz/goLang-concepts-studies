package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello world!", channel)

	fmt.Println("after func write exec")

	for message := range channel {
		fmt.Println(message)
	}
	fmt.Println("End!")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
