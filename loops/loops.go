package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("val: ", i)
	}

	fruits := [...]string{"apple", "orange", "mango", "peach"}

	for index, val := range fruits {
		fmt.Println("val and index", index, val)
	}
}
