package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")
	content := "This needs to be go in a file"

	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("file data: ", string(dataByte))
}
