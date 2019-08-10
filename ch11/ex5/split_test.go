// Extend TestSplit to use a table of inputs and expected outputs.
package ex5

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		str, sep string
		want     int
	}{
		{"a:b:c", ":", 3},
		{"", ":", 1},
		{"a:b:c", "d", 1},
		{"a:b:c", "a:", 2},
		{"abc", "abc", 2},
	}

	for _, test := range tests {
		words := strings.Split(test.str, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.str, test.sep, got, test.want)
		}
	}

}
