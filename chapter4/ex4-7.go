package chapter4

import (
	"unicode/utf8"
)

// Modify reverse to reverse the characters of a []byte slice that represents
// a UTF-8-encoded string, in place. Can you do it without allocating new memory?

func rev(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s []byte) {
	for i := 0; i < len(s); i++ {
		_, size := utf8.DecodeRune(s[i:])
		if size == 1 {
			continue
		}

		rev(s[i : i+size])
		i += size - 1
	}

	rev(s)
}
