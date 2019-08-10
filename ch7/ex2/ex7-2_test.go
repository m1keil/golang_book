package ch7

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEx7_2(t *testing.T) {
	w, c := CountingWriter(&bytes.Buffer{})
	fmt.Fprint(w, "hello")

	if *c != 5 {
		t.Errorf("expected 5 got %v", *c)
	}

	fmt.Fprint(w, "world")
	if *c != 10 {
		t.Errorf("expected 10 got %v", *c)
	}
}
