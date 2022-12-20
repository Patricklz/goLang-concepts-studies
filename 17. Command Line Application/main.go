package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start point")

	aplication := app.Create()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
