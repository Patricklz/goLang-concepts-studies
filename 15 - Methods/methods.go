package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user data of: %s on database\n", u.name)
}

func (u user) isLegalAge() bool {
	return u.age >= 18
}

func (u *user) doBirthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 20}
	user1.save()

	user2 := user{"User 2", 30}
	user2.save()

	legalAge := user2.isLegalAge()
	fmt.Println(legalAge)

	user2.doBirthday()
	fmt.Println(user2.age)
}
