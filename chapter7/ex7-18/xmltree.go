package main

/*
 Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of
 generic nodes that represents it. Nodes are of two kinds: CharData nodes represent text strings, and Element nodes
 represent named elements and their attributes. Each element node has a slice of child nodes.
*/

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	return fmt.Sprintf("%v: %v,", n.Type.Local, n.Children)
}

func nodes(decoder *xml.Decoder) (root Node) {
	var stack []*Element
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, e)
			}
			stack = append(stack, e)

		case xml.EndElement:
			if len(stack) != 1 {
				stack = stack[:len(stack)-1]
			} else {
				root = stack[len(stack)-1]
			}

		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}

	return
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	fmt.Println(nodes(dec))
}