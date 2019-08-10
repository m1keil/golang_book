// Write a function to populate a mapping from element names—p, div, span, and so on—to the number of elements with that
//name in an HTML document tree.
package ex2

import "golang.org/x/net/html"

func count(sum map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		sum[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(sum, c)
	}
}
