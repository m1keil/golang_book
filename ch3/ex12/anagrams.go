package ex12

import "strings"

func IsAnagram(left, right string) bool {
	if len(left) != len(right) {
		return false
	}

	for _, c := range left {
		if !strings.ContainsRune(right, c) {
			return false
		}
	}

	return true
}
