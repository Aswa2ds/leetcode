package codetop

import "math"

func maxSubArray(nums []int) int {
	left := math.MinInt32
	result := left

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, num := range nums {
		left = max(num, num+left)
		result = max(left, result)
	}
	return result

}
