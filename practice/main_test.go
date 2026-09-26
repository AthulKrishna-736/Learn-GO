package main

import (
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	go func() {
		time.Sleep(time.Second)
	}()

	time.Sleep(2 * time.Second)
}
