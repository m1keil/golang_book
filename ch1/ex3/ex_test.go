// Experiment to measure the difference in running time between our potentially inefficient versions and the one that
// uses strings.Join.

/*
$ go test -bench=. --args hello world
goos: linux
goarch: amd64
pkg: golang/ch1/ex3
BenchmarkEcho1-4        10000000               130 ns/op
BenchmarkEcho2-4        10000000               132 ns/op
BenchmarkEcho3-4        20000000                65.9 ns/op
PASS
ok      golang/ch1/ex3  4.286s
*/

package ex3

import (
	"os"
	"strings"
	"testing"
)

func echo3() {
	strings.Join(os.Args[1:], " ")
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		echo3()
	}
}
