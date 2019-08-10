package ch7

/*
 Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the
 original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer.
*/

import (
	"io"
)

type wrapper struct {
	wrapped io.Writer
	counter int64
}

func (w *wrapper) Write(p []byte) (int, error) {
	written, err := w.wrapped.Write(p)
	w.counter += int64(written)

	return written, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	r := &wrapper{wrapped: w}

	return r, &r.counter
}
