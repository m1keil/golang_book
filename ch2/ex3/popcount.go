// Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var sum int
	for i := uint(0); i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}

	return sum
}
