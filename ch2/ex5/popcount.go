// The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCount that counts bits by using
// this fact, and assess its performance.
package popcount

func PopCount(x uint64) int {
	var c int
	for {
		if x == 0 {
			return c
		}

		x = x & (x - 1)
		c++
	}
}
