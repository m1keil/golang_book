package main

/*
Make the broadcaster announce the current set of clients to each new arrival. This requires that the clients set and
the entering and leaving channels record the client name too.
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//type client chan<- string // an outgoing message channel
type client struct {
	ch		chan string // an outgoing message channel
	name	string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			var everyone []string
			for c := range clients {
				everyone = append(everyone, c.name)
			}
			cli.ch <- fmt.Sprintf("connected (%d): %s", len(clients), strings.Join(everyone, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	client := client{ch: make(chan string), name: conn.RemoteAddr().String()}
	go clientWriter(conn, client.ch)

	who := conn.RemoteAddr().String()
	client.ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
