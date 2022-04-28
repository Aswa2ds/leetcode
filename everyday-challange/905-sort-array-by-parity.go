package everydaychallange

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left != right {
		for right > left && nums[right]%2 == 1 {
			right--
		}
		for right > left && nums[left]%2 == 0 {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
