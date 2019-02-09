package chapter7

/*
 The LimitReader function in the io package accepts an io.Reader r and a number of bytes n, and returns another Reader
 that reads from r but reports an end-of-file condition after n bytes. Implement it.
*/

import (
	"fmt"
	"io"
)

type limitreader struct {
	reader   io.Reader
	reminder int64
}

// this exercise is one of the few in the book with a solution.
// the stdlib solution is better. keeping this one for educational purposes :-)
func (s *limitreader) Read(p []byte) (n int, err error) {
	d := make([]byte, s.reminder)
	if s.reminder > 0 {
		i, err := s.reader.Read(d)
		s.reminder -= int64(i)
		fmt.Println(len(p))
		copy(p, d)
		return i, err
	} else {
		return 0, io.EOF
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitreader{r, n}
}
