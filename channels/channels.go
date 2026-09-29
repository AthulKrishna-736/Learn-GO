package main

import "fmt"

func main() {
	messages := make(chan int)
	buff := make(chan int, 3)
	async := make(chan int, 2) // unbuffered channel are blocking both sender & receiver required where buffered channel are non blocking

	go func() {
		messages <- 10

		for i := range 3 {
			buff <- i
		}

		buff <- 10

		async <- 11

		close(buff)
	}()

	msg := <-messages
	fmt.Println("channel msg: ", msg)

	// fmt.Println("val1: ", <-messages) fatal error: all goroutines are asleep - deadlock!

	for val := range buff { // internally <-buff receive through range
		fmt.Println("buff chan: ", val)
	}

	close(messages)

	fmt.Println("channel after close: ", <-messages)

}
