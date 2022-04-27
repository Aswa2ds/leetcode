package everydaychallange

import "math/rand"

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (this *Solution) Pick(target int) (answer int) {
	count := 0
	for i, num := range *this {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				answer = i
			}
		}
	}
	return
}
