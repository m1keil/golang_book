// Add a method Elems that returns a slice containing the elements of the set, suitable for iterating over with a range
// loop.
package intset

func (s *IntSet) Elems() []uint {
	var out []uint
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				out = append(out, uint(SIZE*i+j))
			}
		}
	}

	return out
}
