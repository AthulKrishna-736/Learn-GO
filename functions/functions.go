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

func variadicFunction(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}

	return sum
}

func main() {
	myMessage()
	hello("athul", 10)
	hello("anand", 22)
	fmt.Println(add(1, 2))

	a, b := add(2, 3)
	fmt.Println("total values: ", a, b)

	res := variadicFunction(1, 2, 3, 4)
	fmt.Printf("\nvariadic 1: %d", res)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	res1 := variadicFunction(nums...)
	fmt.Printf("\nvariadic 2: %d", res1)

}
