package main

import (
	"fmt"
)

var a = 10
var b = 22

func main() {
	if a > b {
		fmt.Println("a greater than b")
	} else {
		fmt.Println("b greater thatn a")
	}

	if b > 200 {
		fmt.Println("b not less than 200")
	} else if b > 20 {
		fmt.Println("b not less than 20")
	} else {
		fmt.Println("b is some other value")
	}

	switch 100 {
	case 1:
		fmt.Println("case matched equal to a")
	case 10:
		fmt.Println("sample ccase matched here 10")
	default:
		fmt.Println("no case matched")
	}
}
