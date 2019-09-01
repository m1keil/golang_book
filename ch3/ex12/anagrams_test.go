// Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters
// in a different order.
package ex12

import "testing"

func TestIsAnagram(t *testing.T) {
	d := []struct {
		left, right string
		expects     bool
	}{
		{"hello", "world", false},
		{"Damon Albarn", "Dan Abnormal", true},
		{"evil", "vile", true},
		{"forty five", "over fifty", true},
		{"hello", "helo", false},
	}

	for _, i := range d {
		if IsAnagram(i.left, i.right) != i.expects {
			t.Errorf("IsAnagram(%v, %v) != %v", i.left, i.right, i.expects)
		}
	}
}
