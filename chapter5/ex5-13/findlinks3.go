package main

/*
 Modify crawl to make local copies of the pages it finds, creating directories
 as necessary. Don’t make copies of pages that come from a different domain.
 For example, if the original page comes from golang.org, save all files from
 there, but exclude ones from vimeo.com.
*/

import (
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	if isInList(os.Args[1:], url) {
		savePage(url)
	}

	return list
}

func savePage(address string) {
	fmt.Println("saving:", address)

	u, err := url.Parse(normalizePath(address))
	if err != nil {
		fmt.Println("unable to parse url")
		return
	}

	dir := path.Join(u.Hostname(), path.Dir(u.EscapedPath()))
	file := path.Base(u.EscapedPath())

	fmt.Printf("%v %v\n", dir, file)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("unable to create dir", dir)
		return
	}

	f, err := os.Create(path.Join(dir, file))
	if err != nil {
		fmt.Println("unable to create file", f.Name())
		return
	}

	resp, err := http.Get(address)
	if err != nil {
		fmt.Println("unable to download", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println("failed to write:", f.Name())
	}

	if err = f.Close(); err != nil {
		fmt.Println("failed to write:", f.Name())
	}
}

// normalize url path:
// http://www.example.com/                      -> http://www.example.com/index.html
// http://www.example.com/abc/                  -> http://www.example.com/abc/index.html
// http://http://www.example.com/abc/index.html -> http://www.example.com/abc/index.html
func normalizePath(address string) string {
	if path.Ext(address) == ".html" {
		return address
	}

	u, _ := url.Parse(address)
	u.Path = path.Join(u.EscapedPath(), "index.html")
	return u.String()
}

// checks whether the hostname of root exists in given list roots
func isInList(roots []string, root string) bool {
	u, err := url.Parse(root)
	if err != nil {
		return false
	}
	for _, i := range roots {
		i, err := url.Parse(i)
		if err != nil {
			return false
		}
		if i.Hostname() == u.Hostname() {
			return true
		}
	}
	return false
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
