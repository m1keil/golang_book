package chapter7

import (
	"bytes"
	"fmt"
)

/*
 Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4) that reveals the sequence of values in the tree
*/

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	values := appendValues([]int{}, t)
	buff := bytes.Buffer{}

	if len(values) > 0 {
		fmt.Fprintf(&buff, "%v", values[0])
	}
	for _, v := range values[1:] {
		fmt.Fprintf(&buff, ", %v", v)
	}

	return buff.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
