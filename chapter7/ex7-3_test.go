package chapter7

import (
	"testing"
)

func TestEx7_3(t *testing.T) {
	var tt *tree
	tt = add(tt, 5)
	tt = add(tt, 3)
	tt = add(tt, 10)
	tt = add(tt, 1)

	if tt.String() != "1, 3, 5, 10" {
		t.Fatalf("not the expected tree: %v", tt.String())
	}
}
