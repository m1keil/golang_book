package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Write a program wordfreq to report the frequency of each word in an input
text file. Call input.Split(bufio.ScanWords) before the first call to Scan to
break the input into words instead of lines.
*/

func main() {
	words := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error whilr reading input: %v\n", err)
		os.Exit(1)
	}

	for k, v := range words {
		fmt.Printf("%-15s %v\n", k, v)
	}
}
