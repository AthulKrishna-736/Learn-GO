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

func closureCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
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

	anonymous := func(a, b, c int) int {
		return a - b - c
	}

	res3 := anonymous(1, 2, 3)
	fmt.Printf("\nAnonymous: %d", res3)

	c := closureCounter()
	d := closureCounter()
	fmt.Printf("\n counter 1: %d", c())
	fmt.Printf("\n counter 1: %d", c())
	fmt.Printf("\n counter 1: %d", c())

	fmt.Printf("\n counter 2: %d", d())
	fmt.Printf("\n counter 2: %d", d())
	fmt.Printf("\n counter 2: %d", d())

}
