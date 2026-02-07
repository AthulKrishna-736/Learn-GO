package main

import (
	"fmt"
)

// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

// var a = make(map[KeyType]ValueType) - using make creating map
// b := make(map[KeyType]ValueType)

var a = map[string]string{"car1": "ford", "car2": "porshe"}

func main() {
	b := map[int]string{1: "sample", 2: "sampl2"}

	c := make(map[string]int)

	c["s1"] = 1
	c["s2"] = 2

	delete(c, "s1")

	fmt.Println("map1: ", a, len(a))
	fmt.Println("map2: ", b, len(b))
	fmt.Println("map3: ", c, len(c))

	for key, val := range a {
		fmt.Println("check: ", key, val)
	}
}
