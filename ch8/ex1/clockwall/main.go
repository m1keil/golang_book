package main

// Modify clock2 to accept a port number, and write a program, clockwall,
// that acts as a client of several clock servers at once, reading the times
// from each one and displaying the results in a table, akin to the wall of
// clocks seen in some business offices. If you have access to geographically
// distributed computers, run instances remotely; otherwise run local
// instances on different ports with fake time zones.

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	servers := map[string]string{}

	args := os.Args[1:]
	for _, a := range args {
		label, addr := splitArg(a)
		servers[label] = addr
	}

	fmt.Printf("%-20s %6s\n", "Region", "Time")
	for label, addr := range servers {
		time, err := readTime(addr)
		if err != nil {
			fmt.Printf("%-20s %6s\n", label, "unable to fetch time")
			continue
		}
		fmt.Printf("%-20s %6s\n", label, time)
	}
}

func readTime(server string) (string, error) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	return scanner.Text(), nil
}

func splitArg(arg string) (label string, addr string) {
	e := strings.Split(arg, "=")
	label = e[0]
	addr = e[1]

	return
}
