package ch5

import (
	"testing"
)

func TestEx5_15(t *testing.T) {

	testcases := []struct {
		numbers []int
		min     int
		max     int
	}{
		{[]int{1, 2, 3}, 1, 3},
		{[]int{0}, 0, 0},
	}

	for _, c := range testcases {
		valMin, err := min(c.numbers...)
		valMax, err := max(c.numbers...)
		if err != nil {
			t.Fatalf("got unexpected err: %v", err)
		}

		if valMin != c.min {
			t.Errorf("expected min=%v, got %v", c.min, valMin)
		}

		if valMax != c.max {
			t.Errorf("expected max=%v, got %v", c.max, valMax)
		}
	}

	_, err := min([]int{}...)
	if err.Error() != "expecting at least 1 number" {
		t.Errorf("unexpcted error: %v", err)
	}

}
