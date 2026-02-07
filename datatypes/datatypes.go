package main

import "fmt"

var num = 10               // infer datatype
var name string = "sample" //type defined variable

var dec = 10.9

var a int //default 0
var b string // default ""
var c bool // default false
const d bool = true // constant varible


func main() {
	msg := "another initialization" // should initialize value while delcaring shorthand type compiler will infer

	fmt.Println(dec)
	fmt.Println(num)
	fmt.Println(name)
	fmt.Println(d)
	fmt.Println(msg)
}
