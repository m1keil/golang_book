package ch7

import "sort"

/*
 The sort.Interface type can be adapted to other uses. Write a function IsPalindrome(s sort.Interface) bool that reports
 whether the sequence s is a palindrome, in other words, reversing the sequence would not change it. Assume that the
 elements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i).
*/


func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1 ; i < s.Len(); i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) { continue }
		return false
	}

	return true
}