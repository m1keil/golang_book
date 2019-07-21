package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// Implement a concurrent File Transfer Protocol (FTP) server. The server
// should interpret commands from each client such as cd to change
// directory, ls to list a directory, get to send the contents of a file, and
// close to close the connection. You can use the standard ftp command as the
// client, or write your own.

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	log.Println("New incoming connection", c.RemoteAddr())
	fmt.Fprintf(c, "%d %v\n", 220, "Service ready")

	userDtpIP := c.RemoteAddr().String()
	userDtpPort := 0

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		incoming := scanner.Text()
		log.Println(c.RemoteAddr(), incoming)
		cmd, arg := parseInput(incoming)

		switch cmd {
		case "USER":
			fmt.Fprintf(c, "%d %v\r\n", 331, "User name ok, need password")

		case "PASS":
			fmt.Fprintf(c, "%d %v\r\n", 230, "User logged in, proceed")

		case "SYST":
			fmt.Fprintf(c, "%d %v\r\n", 215, "Golang")

		case "QUIT":
			fmt.Fprintf(c, "%d %v\r\n", 221, "Service closing control connection")
			continue

		case "PORT":
			userDtpIP, userDtpPort = parsePortArgs(arg)
			fmt.Fprintf(c, "%d %v\r\n", 200, "Command ok")

		default:
			fmt.Fprintf(c, "%d %v\r\n", 502, "Command not implemented")
		}
	}

	if err := c.Close(); err != nil {
		log.Println(err)
	}
	log.Println("Closing connection", c.RemoteAddr())
}

func parseInput(s string) (cmd string, args string) {
	tokens := strings.Split(s, " ")

	cmd = tokens[0]
	if len(tokens) > 1 {
		args = tokens[1]
	}

	return
}

func parsePortArgs(s string) (string, int) {
	tokens := strings.Split(s, ",")

	p1, _ := strconv.Atoi(tokens[4])
	p2, _ := strconv.Atoi(tokens[5])

	port := p1*256 + p2
	host := fmt.Sprintf("%s.%s.%s.%s", tokens[0], tokens[1], tokens[2], tokens[3])

	return host, port
}
