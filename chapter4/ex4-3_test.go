package chapter4

import (
	"testing"
)

func TestEx4_3(t *testing.T) {
	have := [...]int{1, 2, 3, 4, 5}
	want := [...]int{5, 4, 3, 2, 1}

	reverse(&have)
	if have != want {
		t.Errorf("expected %v, got %v", want, have)
	}
}
