package main

import (
	"fmt"
)

// slice_name := []datatype{values} - arr length not needed for this initialization
// slice_name := make([]type, length, capacity) - using make method to create slice

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8} // we can do global slice still not recommended

func main() {
	slice1 := []int{1, 2, 3} // local slices are preferred more better
	slice2 := []string{}
	mySlice := arr[0:5] // slice from array
	mySlice1 := make([]int, 3, 4)
	mySlice2 := make([]string, 4, 5)

	mySlice3 := append(arr, slice1...)

	slice1 = append(slice1, 4)
	slice2 = append(slice2, "a")
	slice2 = append(slice2, "b")
	slice2 = append(slice2, "c")

	mySlice1 = append(mySlice1, 1)
	mySlice2 = append(mySlice2, "s")

	fmt.Println("append slice: ", mySlice3, len(mySlice3), cap(mySlice3))

	fmt.Println("make int slice: ", mySlice1, len(mySlice1), cap(mySlice1))
	fmt.Println("make string slice: ", mySlice2, len(mySlice2), cap(mySlice2))

	fmt.Println("slice from arr: ", mySlice, len(mySlice), cap(mySlice))
	fmt.Println("arr: ", arr, len(arr), cap(arr))
	fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
}
