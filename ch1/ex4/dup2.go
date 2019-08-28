// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, locations)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, locations)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t(%v)\n", n, line, strings.Join(locations[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, locations map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		locations[input.Text()] = append(locations[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
