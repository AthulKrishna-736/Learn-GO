package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	if stop := timer.Stop(); stop {
		fmt.Println("timer stopped before expiring")
	}

	// <-timer.C
	// fmt.Println("timer expired")

	<-time.After(2 * time.Second)
	fmt.Println("alternative timer expired")

	timer1 := time.NewTimer(5 * time.Second)
	if !timer1.Stop() {
		<-timer1.C
	}

	timer1.Reset(2 * time.Second)
	<-timer1.C
	fmt.Println("timer reset successfully")
}
