package chapter5

import "fmt"

/*
 Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments?
 Write variants that require at least one argument.
*/

func max(nums ...int) (max int, err error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("expecting at least 1 number")
	}
	max = nums[0]

	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func min(nums ...int) (min int, err error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("expecting at least 1 number")
	}
	min = nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return
}
