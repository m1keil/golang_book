//Define a variadic (*IntSet).AddAll(...int) method that allows a list of values to be added, such as s.AddAll(1, 2, 3)
package intset

// AddAll ints to the set
func (s *IntSet) AddAll(ints ...int) {
	for _, i := range ints {
		s.Add(i)
	}
}
