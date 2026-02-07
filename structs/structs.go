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

func printPerson(pers Person) {
	fmt.Println("func person: ", pers.name, pers.age, pers.mark, pers.test)

}

func main() {

	var per1 Person
	var per2 Person

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

	printPerson(per1)

}
