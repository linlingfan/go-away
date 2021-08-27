package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	abort := make(chan int, 1)
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		if countdown == 5 {
			abort <- -1
		}
		select {
		case <-tick:
			// Do nothing.
		case v := <-abort:
			fmt.Printf("Launch aborted!value：%+v \n", v)
			return
		case <-done:

		}
	}
	println("launch")
}
