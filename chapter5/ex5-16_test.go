package chapter5

import (
	"testing"
)

func TestEx5_16(t *testing.T) {

	testcases := []struct {
		words    []string
		sep      string
		expected string
	}{
		{[]string{"hello", "world"}, " ", "hello world"},
		{[]string{"hello"}, "&", "hello"},
	}

	for _, c := range testcases {
		val := Join(c.sep, c.words...)
		if val != c.expected {
			t.Errorf("expected: %v\ngot: %v\n", c.expected, val)
		}
	}
}
