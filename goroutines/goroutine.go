package main

import (
	"fmt"
	"time"
)

func sample(str string) {
	for i := range 3 {
		fmt.Printf("\nstr: %s, %d", str, i)
	}
}

func main() {
	fmt.Println("started")

	sample("check 1")

	go sample("check 2")
	time.Sleep(time.Second)
}
