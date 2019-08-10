package ex5

import (
	"testing"
)

func TestEx4_5(t *testing.T) {
	have := []string{"a", "a", "b", "c", "c", "c"}
	want := []string{"a", "b", "c"}

	have = deduplicate(have)
	for i := range have {
		if have[i] != want[i] {
			t.Errorf("error on %v: expected %v, got %v", i, want[i], have[i])
		}
	}

	have = deduplicate(nil)
	if have != nil {
		t.Errorf("expected nil but got %v", have)
	}

	have = deduplicate([]string{})
	if len(have) != 0 {
		t.Errorf("expected empty slice but got %v", have)
	}
}
