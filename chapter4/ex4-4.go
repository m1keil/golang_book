package chapter4

// Write a version of rotate that operates in a single pass.

func rotate(s *[]int) {
	if len(*s) == 0 {
		return
	}

	first := (*s)[0]
	for i := 0; i < len(*s)-1; i++ {
		(*s)[i] = (*s)[i+1]
	}
	(*s)[len(*s)-1] = first
}
