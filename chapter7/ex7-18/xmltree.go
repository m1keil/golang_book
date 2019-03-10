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

func getElements(dec *xml.Decoder) Node {
	fmt.Println("start func")
	var n Node


	tok, err := dec.Token()
	if err == io.EOF {
		fmt.Println("reached end")
		return n
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}

	switch tok := tok.(type) {
	case xml.StartElement:
		fmt.Println("start element", tok.Name)
		//n, _ = n.(Element)
		e := Element{Type: tok.Name, Attr: tok.Attr}
		e.Children = append(e.Children, getElements(dec))
		fmt.Println(e)
	case xml.EndElement:
		fmt.Println("end element")
		return n
	case xml.CharData:
		fmt.Println("txt element")
		return CharData(tok)
	}
	fmt.Println("end function", n)
	return n

}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	root := getElements(dec)
	fmt.Println(root)
}
