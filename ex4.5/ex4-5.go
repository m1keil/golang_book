package main

import (
	"fmt"
)

// Write an in-place function to eliminate adjacent duplicates in a []string slice

func main() {
	s := []string{"x", "y", "y", "y", "w", "a", "a"}

	s = deduplicate(s)
	fmt.Println(s)

}

func deduplicate(s []string) []string {
	removed := 0
	for i := 0; i < len(s)-1-removed; {
		if s[i] != s[i+1] {
			i++
			continue
		}

		copy(s[i:], s[i+1:])
		removed++
	}

	return s[:len(s)-removed]
}
