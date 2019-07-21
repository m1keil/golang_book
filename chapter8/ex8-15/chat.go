package main

/*
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck.
Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it.
Alternatively, add buffering to each client’s outgoing message channel so that most messages are not dropped; the
broadcaster should use a non-blocking send to this channel.
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
	IP		net.Addr
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
				select {
				case client.C <- msg:
				// if client channel not ready, skip
				default:
				}
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

func getUsername(conn net.Conn) (string, error) {
	_, err := fmt.Fprintln(conn, "Welcome! please choose a username.")
	if err != nil {
		return "", err
	}
	var user string
	_, err = fmt.Fscanln(conn, &user)
	if err !=nil {
		return "", err
	}

	return user, nil
}

func handleConn(conn net.Conn) {
	user, _ := getUsername(conn)
	client := client{C: make(chan string), Name: user, IP: conn.RemoteAddr()}
	go clientWriter(conn, client.C)

	client.C <- "You are " + client.Name
	messages <- client.Name + " has arrived"
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
			messages <- client.Name + ": " + msg
			timeout.Reset(TIMEOUT)
		case <- timeout.C:
			break loop
		}
	}

	leaving <- client
	messages <- client.Name + " has left"
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
