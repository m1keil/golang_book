package ch7

/*
 Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful.
*/

import (
	"bufio"
	"bytes"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}

	return int(*c), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*c += 1
	}

	return int(*c), nil
}
