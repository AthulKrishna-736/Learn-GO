package main

import "fmt"

// Defer - when used create a stack and insert each defer statments in stack and then executed Lifo order of stack
// Each function maintains its own stack of deferred function calls. When the function returns, deferred calls are executed in LIFO order.

func main() {
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/*
Code execution

main function called
main function stack - [world, one, two]

myDefer function called
myDefer function stack - [0,1,2,3,4]

priting
-------
Hello
43210
two
one
world
*/
