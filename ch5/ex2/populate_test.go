package ex2

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestEx5_2(t *testing.T) {
	content := `
	<HTML>
	<HEAD>
	<TITLE>Your Title Here</TITLE>
	</HEAD>
	<BODY BGCOLOR="FFFFFF">
	<CENTER><IMG SRC="clouds.jpg" ALIGN="BOTTOM"> </CENTER>
	<HR>
	<a href="http://somegreatsite.com">Link Name</a>
	is a link to another nifty site
	<H1>This is a Header</H1>
	<H2>This is a Medium Header</H2>
	Send me mail at <a href="mailto:support@yourcompany.com">
	support@yourcompany.com</a>.
	<P> This is a new paragraph!
	<P> <B>This is a new paragraph!</B>
	<BR> <B><I>This is a new sentence without a paragraph break, in bold italics.</I></B>
	<HR>
	</BODY>
	</HTML>
	`

	root, err := html.Parse(strings.NewReader(content))
	if err != nil {
		t.Errorf("unable to parse html: %v", err)
	}

	have := make(map[string]int)
	count(have, root)

	want := map[string]int{
		"html":   1,
		"head":   1,
		"title":  1,
		"body":   1,
		"p":      2,
		"hr":     2,
		"br":     1,
		"h1":     1,
		"h2":     1,
		"b":      2,
		"i":      1,
		"img":    1,
		"center": 1,
		"a":      2,
	}

	for k, v := range have {
		if want[k] != v {
			t.Errorf("expected %v \"%v\" nodes, got %v", want[k], k, v)
		}
	}
	if len(have) != len(want) {
		t.Errorf("want length %v, got %v", len(want), len(have))
	}
}
