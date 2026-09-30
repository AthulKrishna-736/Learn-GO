package main

import (
	"fmt"
	"time"
)

func ping(ch chan<- int) {
	time.Sleep(1 * time.Second)

	ch <- 10
}

func main() {
	a := make(chan int)

	go ping(a)

	select {
	case a1 := <-a:
		fmt.Println("value from channel: ", a1)
	case <-time.After(3 * time.Second):
		fmt.Println("function ping() have timedout")
	case <-time.Tick(time.Second):
		fmt.Println("tick...")
	}
}
