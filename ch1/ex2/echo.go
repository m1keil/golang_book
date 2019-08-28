// Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args[:] {
		fmt.Println(i, a)
	}
}
