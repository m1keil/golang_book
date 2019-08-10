package ex4

import (
	"testing"
)

func TestEx4_4(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4, 5, 1}

	rotate(&have)
	for i := range have {
		if have[i] != want[i] {
			t.Errorf("error on %v: expected %v, got %v", i, want[i], have[i])
		}
	}

	have = nil
	rotate(&have)
	if have != nil {
		t.Error("expected an empty slice")
	}

}
