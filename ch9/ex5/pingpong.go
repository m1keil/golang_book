package main

import (
	"fmt"
	"time"
)

/*
Write a program with two goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?
*/

func main() {
	left := make(chan struct{})
	right := make(chan struct{})
	counter := 0

	// left
	go func() {
		for {
			<-left
			counter += 1
			right <- struct{}{}
		}
	}()

	// right
	go func() {
		for {
			<-right
			counter += 1
			left <- struct{}{}
		}
	}()

	// start
	left <- struct{}{}

	// print stats
	last := 0
	for {
		count := counter
		time.Sleep(time.Second * 1)
		fmt.Println("message/sec:", counter-last)
		last = count
	}
}
