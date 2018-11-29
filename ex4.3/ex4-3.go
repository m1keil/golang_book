package main

import (
	"fmt"
)

// Rewrite reverse to use an array pointer instead of a slice.

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[5]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
