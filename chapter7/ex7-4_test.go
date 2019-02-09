package chapter7

import (
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func TestEx7_4(t *testing.T) {
	htmldoc := `<html>
<body>
<p>hello world</p>
</body>
</html>`

	r := NewReader(htmldoc)
	_, err := html.Parse(r)
	if err != nil {
		t.Fatal("error parsing html:", err)
	}
}
