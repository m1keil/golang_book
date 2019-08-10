// Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to
// print the SHA384 or SHA512 hash instead.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"os"
)

func main() {
	shatype := flag.String("sha", "256", "hash type")
	flag.Parse()

	switch *shatype {
	case "256":
		fmt.Printf("%x\n", compute(sha256.New()))
	case "384":
		fmt.Printf("%x\n", compute(sha512.New384()))
	case "512":
		fmt.Printf("%x\n", compute(sha512.New()))
	default:
		fmt.Println("only supports: 256, 384, 512")
	}
}

func compute(h hash.Hash) []byte {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		_, err := h.Write(input.Bytes())
		if err != nil {
			panic(err)
		}
	}
	return h.Sum([]byte{})
}
