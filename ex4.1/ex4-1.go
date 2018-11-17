package main

// Write a function that counts the number of bits that are different in
// two SHA256 hashes. (See PopCount from Section 2.6.2.)

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	left := sha256.Sum256([]byte("x"))
	right := sha256.Sum256([]byte("X"))

	fmt.Printf("%08b\n", left)
	fmt.Printf("%08b\n", right)

	fmt.Printf("diff: %v", diffBits(left, right))
}

//

func diffBits(left, right [32]byte) int {
	var count int
	for i := range left {
		for j := uint(0); j < 8; j++ {
			if left[i]&(1<<j)^right[i]&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}
