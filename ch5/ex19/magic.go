// Use panic and recover to write a function that contains no return statement yet returns a non-zero value.
package ex19

func magic() (exit int) {
	defer func() {
		if p := recover(); p != nil {
			exit = 1
		}
	}()

	panic("boo!")
}
