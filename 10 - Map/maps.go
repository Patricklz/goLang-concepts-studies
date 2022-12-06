package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Peter",
		"lastname": "Parker",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first":    "Tony",
			"lastname": "Stark",
		},
		"avanger": {
			"power": "super intelligence",
		},
	}

	fmt.Println(user2)
	delete(user2, "avanger")
	fmt.Println(user2)

	user2["power"] = map[string]string{
		"type": "money",
	}

	fmt.Println(user2)

}
