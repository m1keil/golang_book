package main

/*
Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes.
Hint: calling conn.Close() in another goroutine unblocks active Read calls such as the one done by input.Scan().
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//type client chan<- string // an outgoing message channel
type client struct {
	C		chan string // an outgoing message channel
	Name	string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	TIMEOUT = time.Minute * 5 // idle timeout
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for client := range clients {
				client.C <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.C <- listMembers(clients)

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.C)
		}
	}
}

func listMembers(clients map[client]bool) string {
	var everyone []string
	for c := range clients {
		everyone = append(everyone, c.Name)
	}

	return fmt.Sprintf("connected (%d): %s", len(clients), strings.Join(everyone, ", "))

}

func handleConn(conn net.Conn) {
	client := client{C: make(chan string), Name: conn.RemoteAddr().String()}
	go clientWriter(conn, client.C)

	who := conn.RemoteAddr().String()
	client.C <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	timeout := time.NewTimer(TIMEOUT)

	input := make(chan string)
	go clientReader(conn, input)

	loop:
	for {
		select {
		case msg, ok := <- input:
			if !ok {
				break loop
			}
			messages <- who + ": " + msg
			timeout.Reset(TIMEOUT)
		case <- timeout.C:
			break loop
		}
	}

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	close(ch)
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
