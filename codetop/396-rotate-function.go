package codetop

func maxRotateFunction(nums []int) int {
	sum := 0
	rotatedSum := 0
	for i, num := range nums {
		sum += num
		rotatedSum += i * num
	}

	maxRotateSum := rotatedSum
	for i := len(nums) - 1; i > 0; i-- {
		rotatedSum = rotatedSum + sum - len(nums)*nums[i]
		if rotatedSum > maxRotateSum {
			maxRotateSum = rotatedSum
		}
	}

	return maxRotateSum
}
