// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
// Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var b bytes.Buffer

	var start int
	end := len(s)

	if strings.ContainsAny(s[0:1], "+-") {
		start = 1
		b.WriteByte(s[0])
	}

	if i := strings.Index(s, "."); i != -1 {
		end = i
	}

	for i, c := range s[start:end] {
		if i > 0 && (end-start-i)%3 == 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	}

	b.WriteString(s[end:])

	return b.String()
}
