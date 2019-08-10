// Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like
// unicode.IsLetter.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	var letter, number, punct, space int
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)

		switch {
		case unicode.IsLetter(c):
			letter += n
		case unicode.IsNumber(c):
			number += n
		case unicode.IsPunct(c):
			punct += n
		case unicode.IsSpace(c):
			space += n
		}
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	if letter > 0 {
		fmt.Printf("\n%d UTF-8 letters\n", letter)
	}
	if number > 0 {
		fmt.Printf("%d UTF-8 numbers\n", number)
	}
	if punct > 0 {
		fmt.Printf("%d UTF-8 punctuations\n", punct)
	}
	if space > 0 {
		fmt.Printf("%d UTF-8 spaces\n", space)
	}
}
