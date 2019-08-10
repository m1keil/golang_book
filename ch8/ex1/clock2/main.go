package main

// Modify clock2 to accept a port number, and write a program, clockwall,
// that acts as a client of several clock servers at once, reading the times
// from each one and displaying the results in a table, akin to the wall of
// clocks seen in some business offices. If you have access to geographically
// distributed computers, run instances remotely; otherwise run local
// instances on different ports with fake time zones.

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port = flag.Int("port", 8000, "port number")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
