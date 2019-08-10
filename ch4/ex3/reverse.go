package ex3

// Rewrite reverse to use an array pointer instead of a slice.

func reverse(a *[5]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
