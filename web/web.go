package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("sample web requests")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("res: ", res)
	defer res.Body.Close() // should close res as callers responsibility
}
