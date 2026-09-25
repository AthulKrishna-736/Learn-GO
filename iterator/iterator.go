package main

import (
	"fmt"
	"iter"
)

type Iterator struct {
	current int
}

func (I *Iterator) Next() (int, bool) {
	if I.current > 5 {
		return I.current, false
	}

	value := I.current
	I.current++

	return value, true
}

func Numbers() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i < 5; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	c := Iterator{}

	for {
		n, ok := c.Next()
		if !ok {
			break
		}

		fmt.Println(n)
	}

	for n := range Numbers() {
		fmt.Println(n)
	}
}
