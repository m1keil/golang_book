package ex6

import (
	"bytes"
	"testing"
)

func TestEx4_6(t *testing.T) {
	have := []byte{'h', 'e', 'l', 'l', 'o'}
	want := []byte{'h', 'e', 'l', 'l', 'o'}

	have = squash(have)
	if !bytes.Equal(have, want) {
		t.Errorf("expected %v, got %v", want, have)
	}

	have = []byte{'h', '	', '	', 'l', 'o'}
	want = []byte{'h', ' ', 'l', 'o'}

	have = squash(have)
	if !bytes.Equal(have, want) {
		t.Errorf("expected %v, got %v", want, have)
	}
}
