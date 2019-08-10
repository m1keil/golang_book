/* Modify forEachNode so that the pre and post functions return a boolean result indicating whether to continue the
traversal. Use it to write a function ElementByID with the following signature that finds the first HTML element with
the specified id attribute. The function should stop the traversal as soon as a match is found. */
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		findattr(url)
	}
}

func findattr(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(ElementByID(doc, "lowframe"))

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && post(n) {
		return
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var match *html.Node

	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, v := range n.Attr {
				if v.Key == "id" && v.Val == id {
					match = n
					return true
				}
			}
		}
		return false
	}

	forEachNode(doc, pre, nil)

	return match
}
