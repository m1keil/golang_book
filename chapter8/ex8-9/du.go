package main

/* Write a version of du that computes and periodically displays separate totals for each of the root directories. */

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)


var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	totalSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go calcRoot(root, &n, totalSizes)
	}
	go func() {
		n.Wait()
		close(totalSizes)
	}()

	// TODO: read from totalSize and print aggregated size
	var nfiles, nbytes int64
	for size := range totalSizes {
		nbytes += size
		nfiles++
	}

	printDiskUsage(nfiles,nbytes)
}

func calcRoot(dir string, d *sync.WaitGroup, totalSizes chan<- int64) {
	fileSizes := make(chan int64)
	var n sync.WaitGroup

	n.Add(1)
	go walkDir(dir, &n, fileSizes)

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
			totalSizes <- size
		case <-tick:
			printDirUsage(dir, nfiles, nbytes)
		}
	}

	d.Done()
}


func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("total: %d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func printDirUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files  %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
