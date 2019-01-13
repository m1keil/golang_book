package chapter7

import "testing"

func TestEx7_1(t *testing.T) {
	testdata := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	var w WordCounter
	var l LineCounter

	w.Write([]byte(testdata))
	l.Write([]byte(testdata))

	if w != 69 {
		t.Errorf("execpted 69 words, got %v", w)
	}

	if l != 4 {
		t.Errorf("expected 4 lines, got %v", l)
	}
}
