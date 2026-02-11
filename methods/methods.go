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

type Sample struct {
	c int
}

func (s Sample) hello(a, b int) (int, string) {
	return a + b + s.c, "sample output"
}

func main() {
	user1 := User{"sample", 10, true}
	fmt.Println("user1 details: ", user1, user1.Active, user1.age, user1.name)
	user1.greet()

	sample1 := Sample{10}

	res, str := sample1.hello(19, 2)
	fmt.Println("sample res: ", res, str)
}
