// Write tests for the charcount program in Section 4.3
package main

import (
	"strings"
	"testing"
)


func TestCharcount(t *testing.T) {
	tests := []struct{
		input string
		runeMap map[rune]int
		invalid int
	}{
		{"hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o':1}, 0},
		{"\xfe\xff", map[rune]int{}, 2},
		{"שלום", map[rune]int{'ש': 1, 'ל': 1, 'ו': 1, 'ם':1}, 0},
	}


	for _, test := range tests {
		r,_, i, _ := Count(strings.NewReader(test.input))
		if !compare (r, test.runeMap)  {
			t.Errorf("Count(%v) incorrect map %v", test.input, r)
		}
		if i != test.invalid {
			t.Errorf("Count(%v), expected length %v, got %v", test.input, test.invalid, i)
		}
	}
}

func compare (left, right map[rune]int) bool {
	if len(left) != len(right) {
		return false
	}
	for k, v := range left {
		if v != right[k] {
			return false
		}
	}
	return true
}