// Write a variadic version of strings.Join.
package ex16

import "fmt"

func Join(sep string, a ...string) string {
	out := ""
	if len(a) > 0 {
		out = a[0]
	}

	for _, s := range a[1:] {
		out = fmt.Sprintf("%v%v%v", out, sep, s)
	}

	return out
}
