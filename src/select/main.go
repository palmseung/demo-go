package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)
	timerChan := time.After(10 * time.Second)
	tickTimer := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-tickTimer:
			fmt.Println("Tick!")
		case <-timerChan:
			fmt.Println("Timeout")
			return
		}
	}
}
