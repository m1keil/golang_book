package main

/*
 Extend xmlselect so that elements may be selected not just by name, but by their attributes too, in the manner of CSS,
 so that, for instance, an element like <div id="page" class="wide"> could be selected by a matching id or class as well
 as its name.
*/

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type element struct {
	name, id, class string
}

func (e element) String() string {
	return fmt.Sprintf("<%s id=%s class=%s>", e.name, e.id, e.class)
}


func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []element // stack of elements
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			id, class := getIdClass(tok.Attr)
			stack = append(stack, element{tok.Name.Local, id, class}) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stack, tok)
			}
		}
	}
}

// returns the "id" and "class" attribute if defined
// if not, returns empty string
func getIdClass(attrs []xml.Attr) (id, class string) {
	for _, i := range attrs {
		if strings.ToLower(i.Name.Local) == "id" {
			id = i.Value
		}
		if strings.ToLower(i.Name.Local) == "class" {
			class = i.Value
		}
	}

	return
}

// containsAll reports whether x contains the elements of y, in order.
// y is a list of strings. strings can include dot character. first item
// after dot is an id or class attribute.
// i.e div.name can match either <div id="name"> or <div class="name">
func containsAll(x []element, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		items := strings.Split(y[0], ".")
		name := items[0]
		var attr string
		if len(items) > 1 { attr = items[1] }

		if x[0].name == name {
			if attr == "" {
				y = y[1:]
			} else if x[0].class == attr || x[0].id == attr {
				y = y[1:]
			}
		}
		x = x[1:]
	}
	return false
}
