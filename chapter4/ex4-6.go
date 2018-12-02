package chapter4

import (
	"unicode"
)

// Write an in-place function that squashes each rune of adjacent
// Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice
// into a single ASCII space.

func squash(b []byte) []byte {
	out := b[:0]
	flag := false
	for i := range b {
		if unicode.IsSpace(rune(b[i])) {
			flag = true
			continue
		}

		if flag {
			out = append(out, ' ')
			flag = false
		}

		out = append(out, b[i])
	}

	return out
}
