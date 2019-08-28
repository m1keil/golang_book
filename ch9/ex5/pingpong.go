package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
Write a program with two goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?
*/

func main() {
	left := make(chan struct{})
	right := make(chan struct{})
	var counter, last uint64

	// left
	go func() {
		for {
			<-left
			atomic.AddUint64(&counter, 1)
			right <- struct{}{}
		}
	}()

	// right
	go func() {
		for {
			<-right
			atomic.AddUint64(&counter, 1)
			left <- struct{}{}
		}
	}()

	// start
	left <- struct{}{}

	// print stats
	for {
		count := atomic.LoadUint64(&counter)
		time.Sleep(time.Second * 1)
		fmt.Println("message/sec:", count-last)
		last = count
	}
}
