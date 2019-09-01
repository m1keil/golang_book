// Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the
// elements that match one of those names.
package main

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var out []*html.Node
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				out = append(out, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		out = append(out, ElementsByTagName(c, name...)...)
	}

	return out
}
