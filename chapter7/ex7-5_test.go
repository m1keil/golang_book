package chapter7

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestEx7_5(t *testing.T) {
	testdata := "hello world"

	r := LimitReader(strings.NewReader(testdata), 5)

	b, _ := ioutil.ReadAll(r)

	if len(b) != 5 {
		t.Errorf("expected to read 5 bytes, got %v", len(b))
	}
	if string(b) != "hello" {
		t.Errorf("expected to read hello, got %v", b)
	}
}
