package main

import "fmt"

func main() {
	messages := make(chan int)

	go func() {
		messages <- 10
	}()

	msg := <-messages
	fmt.Println("channel msg: ", msg)

	close(messages)

	fmt.Println("channel after close: ", <-messages)
}
