package main

/* Using a select statement, add a timeout to the echo server from Section 8.3
so that it disconnects any client that shouts nothing within 10 seconds */

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	fmt.Println("connected", c.RemoteAddr())
	incoming := make(chan string)
	timeout := time.NewTimer(time.Second * 10)

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			t := input.Text()
			incoming <- t
		}
		close(incoming)
	}()

outer:
	for {
		select {
		case <-timeout.C:
			fmt.Println("timeout reached", c.RemoteAddr())
			break outer

		case text, ok := <-incoming:
			if !ok {
				fmt.Println("reached EOF", c.RemoteAddr())
				break outer
			}
			timeout.Reset(time.Second * 10)
			go echo(c, text, 1*time.Second)
		}
	}

	fmt.Println("closing", c.RemoteAddr())
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
