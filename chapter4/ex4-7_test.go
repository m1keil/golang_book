package chapter4

import (
	"testing"
)

func TestEx4_7(t *testing.T) {
	have := []byte("hello")
	want := []byte("olleh")

	reverse2(have)
	for i := range have {
		if have[i] != want[i] {
			t.Errorf("expected %v, got %v", want, have)
		}
	}

	have = []byte("שלום")
	want = []byte("םולש")

	reverse2(have)
	for i := range have {
		if have[i] != want[i] {
			t.Errorf("expected %v, got %v", want, have)
		}
	}

	have = []byte("")
	want = []byte("")

	reverse2(have)
	for i := range have {
		if have[i] != want[i] {
			t.Errorf("expected %v, got %v", want, have)
		}
	}
}
