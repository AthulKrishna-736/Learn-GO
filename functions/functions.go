package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func hello(name string, age int) {
	fmt.Println("hello", name, age)
}

func add(x int, y int) (res int, text string) {
	res = x + y
	text = "sample"
	return
}

func main() {
	myMessage()
	hello("athul", 10)
	hello("anand", 22)
	fmt.Println(add(1, 2))

	a, b := add(2, 3)
	fmt.Println("total values: ", a, b)

}
