package main

import (
	"fmt"
)

var a = 10
var b = 20

func main() {
	var c = 10 + 30
	var d = a + b

	e := a - b
	f := a / b
	g := a % b

	f += 20

	d++

	fmt.Println("add: ", c)
	fmt.Println("sum: ", d)
	fmt.Println("sub: ", e)
	fmt.Println("div: ", f)
	fmt.Println("mod: ", g)
}
