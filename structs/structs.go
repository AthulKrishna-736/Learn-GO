package main

import (
	"fmt"
)

// type struct_name struct {
//   member1 datatype;
//   member2 datatype;
//   member3 datatype;
//   ...
// }

type Person struct {
	name string
	age  int
	mark float32
	test bool
}

type User struct {
	name        string
	age         uint
	title       *string
	description *string
	marks       *float32
}

func printPerson(pers Person) {
	fmt.Println("func person: ", pers.name, pers.age, pers.mark, pers.test)

}

func main() {

	var per1 Person
	var per2 Person
	var user1 User

	per1.name = "athul"
	per1.age = 10
	per1.mark = 123.3
	per1.test = false

	per2.name = "sample"
	per2.age = 12
	per2.mark = 23.4
	per2.test = true

	fmt.Println("user 1", per1.name, per1.age, per1.mark, per1.test)
	fmt.Println("user 2", per2.name, per2.age, per2.mark, per2.test)

	user1.name = "check"
	user1.age = 19

	if user1.title != nil && *user1.title != "" { // nil pointer dereference because no value so default type set to nil to solve use nil check before type check
		fmt.Println("title is empty string")
	}
	if *user1.marks != 0 {
		fmt.Println("marke are not zero")
	}

	fmt.Println("print struct user: ", user1)

	printPerson(per1)

}
