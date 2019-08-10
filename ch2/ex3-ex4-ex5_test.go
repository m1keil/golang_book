// performance benchmark to all of the PopCount versions
/*

Benchmark results:

BenchmarkPopCountEx3Max-4     	100000000	        21.2 ns/op
BenchmarkPopCountEx3Zero-4    	100000000	        21.6 ns/op
BenchmarkPopCountEx4Max-4     	30000000	        50.1 ns/op
BenchmarkPopCountEx4Zero-4    	30000000	        51.2 ns/op
BenchmarkPopCountEx5Max-4     	30000000	        52.9 ns/op
BenchmarkPopCountEx5Zero-4    	1000000000	         2.63 ns/op
BenchmarkPopCountOrigMax-4    	2000000000	         0.38 ns/op
BenchmarkPopCountOrigZero-4   	2000000000	         0.76 ns/op

Original version is still faster than everything else tested.
Last version (Ex5) is faster with low amount of 1's but getting slower as more 1's added.
*/
package ch2

import (
	"math"
	"testing"
)
import ex3 "golang/ch2/ex3"
import ex4 "golang/ch2/ex4"
import ex5 "golang/ch2/ex5"
import book "gopl.io/ch2/popcount"

func BenchmarkPopCountEx3Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex3.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCountEx3Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex3.PopCount(0)
	}
}

func BenchmarkPopCountEx4Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex4.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCountEx4Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex4.PopCount(0)
	}
}

func BenchmarkPopCountEx5Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex5.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCountEx5Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex5.PopCount(0)
	}
}

func BenchmarkPopCountOrigMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		book.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCountOrigZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		book.PopCount(0)
	}
}

func TestPopCount(t *testing.T) {
	tests := []struct {
		have uint64
		want int
	}{
		{0, 0},
		{math.MaxUint16, 16},
		{math.MaxUint32, 32},
		{math.MaxUint64, 64},
	}

	for _, test := range tests {
		count := ex3.PopCount(test.have)
		if count != test.want {
			t.Errorf("ex3.PopCount(%v) = %v, expected %v", test.have, count, test.want)
		}
	}

	for _, test := range tests {
		count := ex4.PopCount(test.have)
		if count != test.want {
			t.Errorf("ex4.PopCount(%v) = %v, expected %v", test.have, count, test.want)
		}
	}

	for _, test := range tests {
		count := ex4.PopCount(test.have)
		if count != test.want {
			t.Errorf("ex5.PopCount(%v) = %v, expected %v", test.have, count, test.want)
		}
	}
}
