// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"golang/chapter9/ex9-3"
)

func httpGetBody(url string, cancel chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, cancel chan struct{}) (interface{}, error)
}


func Sequential(t *testing.T, m M) {
	var cancel chan struct{}
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, cancel)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}


func Concurrent(t *testing.T, m M) {
	var cancel chan struct{}
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func Cancellation(t *testing.T, m M) {
	for url := range incomingURLs() {
		cancel := make(chan struct{})
		start := time.Now()
		if url == "https://golang.org" {
			go func() {
				time.Sleep(time.Millisecond * 50)
				close(cancel)
			}()
		}
		_, err := m.Get(url, cancel)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %s\n",
			url, time.Since(start), err)
	}
}



func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	Concurrent(t, m)
}

func TestCancellation(t *testing.T) {
	m := memo.New(httpGetBody)
	Cancellation(t, m)
}