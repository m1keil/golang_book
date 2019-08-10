package ch5

import "testing"

func TestEx5_9(t *testing.T) {
	template := "hello $foo, this is your $foo"

	f := func(x string) string {
		return x
	}

	have := expand2(template, f)
	want := "hello foo, this is your foo"

	if have != want {
		t.Fatalf("%v != %v", want, have)
	}
}
