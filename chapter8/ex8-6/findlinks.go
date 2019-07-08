package main

/* Add depth-limiting to the concurrent crawler. That is, if the user sets
-depth=3, then only URLs reachable by at most three links will be fetched. */

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) []string {
	fmt.Println(depth, url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

// Worklist is a combination of a list and its depth
type Worklist struct {
	list  []string
	depth int
}

func main() {
	depth := flag.Int("depth", 3, "links depth")
	flag.Parse()

	worklist := make(chan Worklist)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- Worklist{flag.Args(), 1} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		w := <-worklist
		for _, link := range w.list {
			if !seen[link] && w.depth <= *depth {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- Worklist{crawl(link, w.depth), w.depth + 1}
				}(link)
			}
		}
	}
}
