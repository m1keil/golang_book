package chapter7

import (
	"sort"
	"testing"
)

func TestEx7_10(t *testing.T) {
	x := sort.StringSlice{"a","b", "c", "b", "a"}
	y := sort.StringSlice{"a","b", "c"}

	if !IsPalindrome(x) {
		t.Fatal("failure to detect polindrome")
	}
	if IsPalindrome(y) {
		t.Fatal("failure to detect not a polindrome")
	}
}

