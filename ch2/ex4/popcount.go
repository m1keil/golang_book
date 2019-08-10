// Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the rightmost
// bit each time. Compare its performance to the table-lookup version.
package popcount

func PopCount(x uint64) int {
	var sum int
	for i := uint(0); i < 64; i++ {
		sum += int(x & 1)
		x = x >> 1
	}

	return sum
}
