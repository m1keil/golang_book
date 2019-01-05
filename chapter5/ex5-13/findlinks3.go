package main

/*
 Modify crawl to make local copies of the pages it finds, creating directories
 as necessary. Don’t make copies of pages that come from a different domain.
 For example, if the original page comes from golang.org, save all files from
 there, but exclude ones from vimeo.com.
*/

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"./links"
	"golang.org/x/net/html"
)

func breadthFirst(f func(string, bool) []string, worklist []string) {
	seen := make(map[string]bool)
	roots := worklist
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, isInList(roots, item))...)
			}
		}
	}
}

func crawl(url string, save bool) []string {
	fmt.Println(url)
	list, doc, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	if save {
		savePage(url, doc)
	}

	return list
}

func savePage(u string, page *html.Node) {
	fmt.Println("saving:", u)

	dir, file := normalizePath(url)
	fmt.Printf("dir %v file %v\n", dir, file)
	err := os.Mkdir(dir, 0755)
	if !os.IsExist(err) {
		fmt.Println("unable to create dir", dir)
		return
	}

	f, err := os.Create(filepath)
	if os.IsExist(err) {
		fmt.Println("unable to create file", filepath)
		return
	}
	err = html.Render(f, page)
	if err != nil {
		fmt.Println("failed to write:", filepath)
	}
}

func normalizePath(u string) (string, string) {
	url, err := url.Parse(u)

}

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
