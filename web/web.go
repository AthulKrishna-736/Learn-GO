package main

import (
	"encoding/json"
	"fmt"
)

const url = "https://lco.dev"

type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Status bool     `json:"status"`
	Fruits []string `json:"fruits"`
}

func main() {
	// fmt.Println("sample web requests")

	// res, err := http.Get(url)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("res: ", res)
	// defer res.Body.Close() // should close res as callers responsibility

	a := User{"Sample", 10, false, []string{"apple", "mango", "banana"}}

	res, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("json: ", string(res))

	var b User

	err1 := json.Unmarshal(res, &b)

	if err1 != nil {
		fmt.Println(err1.Error())
	}

	fmt.Println("Struct: ", b)

	isValid := json.Valid(res)
	fmt.Println("is proper json: ", isValid)
}
