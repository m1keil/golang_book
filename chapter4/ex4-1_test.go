package chapter4

import (
	"crypto/sha256"
	"testing"
)

func TestEx4_1(t *testing.T) {
	left := sha256.Sum256([]byte("x"))
	right := sha256.Sum256([]byte("x"))

	if d := diffBits(left, right); d != 0 {
		t.Errorf("Expected 0 but got %v", d)
	}

	right = sha256.Sum256([]byte("X"))
	if d := diffBits(left, right); d != 125 {
		t.Errorf("Expected 125 but got %v", d)
	}

	left = sha256.Sum256([]byte{})
	right = sha256.Sum256([]byte{})
	if d := diffBits(left, right); d != 0 {
		t.Errorf("Expected 0 but got %v", d)
	}
}
