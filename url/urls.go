package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://localhost:3000/level?user=sample&age=20"

func main() {
	fmt.Println("url: ", myUrl)
	res, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("myUrl: ", res.Scheme)
	fmt.Println("myUrl: ", res.Host)
	fmt.Println("myUrl: ", res.Port())
	fmt.Println("myUrl: ", res.Path)
	fmt.Println("myUrl: ", res.RawQuery)
	fmt.Println("myUrl: ", res.Query())

	data := res.Query()

	fmt.Println("query: ", data["user"])
	fmt.Println("query: ", data["age"])

	createUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "level",
		RawQuery: "user=super&age=10",
	}

	for _, val := range data {
		fmt.Println(val)
	}

	fmt.Println("createdUrl: ", createUrl)
}
