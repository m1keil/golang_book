package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

/*
Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that requests several URLs
concurrently. As soon as the first response arrives, cancel the other requests.
*/


func main() {
	cancel := make(chan struct{})
	responses := make(chan []byte, len(os.Args[1:]))

	wg := &sync.WaitGroup{}

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go get(url, responses, cancel, wg)
	}

	first := <-responses
	close(cancel)

	fmt.Fprintf(os.Stdout, "%s", first)
	wg.Wait()
}

func get(url string, responses chan<- []byte, cancel <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading %s: %v\n", url, err)
		return
	}

	responses<-b
}