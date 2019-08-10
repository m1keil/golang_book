// Write a function to print the contents of all text nodes in an HTML document tree.
// Do not descend into <script> or <style> elements, since their contents are not visible in a web browser.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "printtext: %v\n", err)
		os.Exit(1)
	}

	print(doc)
}

func print(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		data := strings.TrimSpace(n.Data)
		if len(data) != 0 {
			fmt.Println(data)
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		print(c)
	}
}
