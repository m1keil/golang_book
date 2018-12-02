package chapter4

// Write a function that counts the number of bits that are different in
// two SHA256 hashes. (See PopCount from Section 2.6.2.)

func diffBits(left, right [32]byte) (diff int) {
	for i := range left {
		for j := uint(0); j < 8; j++ {
			if left[i]&(1<<j)^right[i]&(1<<j) != 0 {
				diff++
			}
		}
	}

	return
}
