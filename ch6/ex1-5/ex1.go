/*
Implement these additional methods:
func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/
package intset

// Len return the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
	copy(words, s.words)
	return &IntSet{words}
}
