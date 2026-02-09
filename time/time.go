package main

import (
	"fmt"
	"time"
)

func statusCheck() string {
	return "status completed"
}

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01 Jan 2006 15:04:05 Monday"))

	createdTime := time.Date(2020, time.December, 19, 19, 19, 19, 19, time.UTC)
	fmt.Println(createdTime.Format("01-02-2006 Monday 15:04:05"))

	fmt.Println(presentTime.Year() > createdTime.Year())

	calc := presentTime.Add(24 * time.Hour)

	fmt.Println("check hour diff: ", presentTime, calc)

	fmt.Println("layout using with time: ", calc.Format(time.Kitchen))

	c := time.Tick(5 * time.Second) // we case use any as per out needs time.Minutes, time.Hour, time.second as per our needs

	fmt.Println("timer working", c)

	for val := range c {
		fmt.Println("time runing...", val, c, statusCheck())
	}

}
