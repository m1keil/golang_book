package ch7

import "io"

/*
 The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading from its
 argument, a string. Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take
 input from a string.
*/

type MyString struct {
	str string
	p   int
}

func (s *MyString) Read(p []byte) (n int, err error) {
	if s.p > len(s.str) {
		return 0, io.EOF
	}
	n = copy(p, s.str)
	s.p += n
	return
}

func NewReader(s string) *MyString {
	m := MyString{s, 0}
	return &m
}
