package main

import "fmt"

func main() {
    channel := make(chan string, 2)
    channel <- "Hello World!"
    channel <- "Programing in Go!"


    message := <- channel

    fmt.Println(message)
}