package main

import "fmt"

type User struct {
	name   string
	age    int
	status bool
}

func (u *User) updateUser(name string, age int) {
	u.name = name
	u.age = age
}

func main() {
	i := 10
	j := 11
	fmt.Println("val: ", i, j, &i, &j)
	p := &i
	fmt.Println("value of p ", p, *p)
	changeValue(p)
	fmt.Println("changed value ", i)

	user1 := User{"sample", 10, false}
	user1.updateUser("new_name", 20)
	fmt.Println("user data: ", user1)}

func changeValue(a *int) {
	*a = 20
}
