package chapter5

import (
	"testing"
)

func TestEx5_19(t *testing.T) {
	val := magic()
	if val == 0 {
		t.Errorf("expected non zero, got %v", val)
	}
}
