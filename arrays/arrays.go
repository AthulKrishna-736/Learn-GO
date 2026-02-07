package main

import (
	"fmt"
)

// var array_name = [length]datatype{values} // here length is defined
// var array_name = [...]datatype{values} // here length is inferred

var nums = [5]int{1, 2, 3, 4, 5}
var list = [...]string{"a", "b", "c", "d", "e"}
var arr1 = [5]int{2}
var arr2 = [5]string{"a"}
var arr4 = [...]int{}

func main() {
	arr3 := [...]float32{12.4, 23.3, 43, .34}
	fmt.Println("number array: ", nums, " get value: ", nums[0])
	fmt.Println("string array: ", list, " get value: ", list[2])
	fmt.Println("float array: ", arr3, " get value: ", arr3[3])

	fmt.Println("unintialized arrays: ", arr1, arr2)

	nums[4] = 33
	list[2] = "sample"

	fmt.Println("updated: ", nums)
	fmt.Println("updated: ", list)

	fmt.Println("arr lengths: ", len(nums), len(arr1), len(arr2), len(arr3), len(arr4))
}
