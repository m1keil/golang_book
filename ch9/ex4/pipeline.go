package main

import (
	"fmt"
	"time"
)

/*
Construct a pipeline that connects an arbitrary number of goroutines with channels. What is the maximum number of
pipeline stages you can create without running out of memory? How long does a value take to transit the entire pipeline?
*/

var MAX = 5000000

func main() {
	start := time.Now()
	head := make(chan struct{})
	length := 0

	go func() {
		// monitor length
		for {
			fmt.Println("length:", length)
			time.Sleep(time.Second * 1)
		}
	}()

	last := head
	for i:=0; i< MAX; i++{
		out := make(chan struct{})
		go func(in, out chan struct{}) {
			<- in
			out <- struct{}{}
		}(last, out)

		length += 1
		last = out
	}
	fmt.Printf("time to init: %s\n", time.Since(start))

	start = time.Now()
	head <- struct{}{}
	<- last
	fmt.Printf("time to run: %s\n", time.Since(start))
}

