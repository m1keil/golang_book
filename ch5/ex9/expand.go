// Write a function expand(s string, f func(string) string) string that replaces each substring “$foo” within s by the
// text returned by f("foo").
package ch5

import (
	"strings"
)

// is this version is cheating?
func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

// same but with recursion?
func expand2(s string, f func(string) string) string {
	r := []rune(s)
	if len(r) < 4 {
		return s
	}

	if string(r[0:4]) == "$foo" {
		return f("foo") + expand2(string(r[4:]), f)
	}

	return string(r[0]) + expand2(string(r[1:]), f)
}
