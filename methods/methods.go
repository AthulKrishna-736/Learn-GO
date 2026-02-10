package main

import "fmt"

type User struct {
	name   string
	age    int
	Active bool
}

func (u User) greet() {
	fmt.Println("Hello ,", u.name)
}

func main() {
	user1 := User{"sample", 10, true}
	fmt.Println("user1 details: ", user1, user1.Active, user1.age, user1.name)
	user1.greet()
}
