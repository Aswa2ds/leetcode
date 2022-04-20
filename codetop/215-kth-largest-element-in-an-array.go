package codetop

import "math"

func findKthLargest(nums []int, k int) int {
	heap := make([]int, 0)
	capacity := k

	var top func(heap []int) int
	var shiftup func(heap []int, j int)
	var shiftdown func(heap []int, i int)

	top = func(heap []int) int {
		if len(heap) == 0 {
			return math.MinInt
		}
		return heap[0]
	}

	shiftup = func(heap []int, j int) {
		if j == 0 {
			return
		}
		i := (j - 1) / 2
		if heap[i] > heap[j] {
			heap[i], heap[j] = heap[j], heap[i]
			shiftup(heap, i)
		}
	}

	shiftdown = func(heap []int, i int) {
		if i >= len(heap)/2 {
			return
		}
		j := i
		j0 := 2*i + 1
		if heap[i] > heap[j0] {
			j = j0
		}
		j1 := j0 + 1
		if j1 < len(heap) && heap[j] > heap[j1] {
			j = j1
		}
		if i == j {
			return
		}
		heap[i], heap[j] = heap[j], heap[i]
		shiftdown(heap, j)
	}

	push := func(heap *[]int, num int) {
		if len(*heap) == capacity && num < top(*heap) {
			return
		}
		if len(*heap) == capacity {
			(*heap)[0] = num
			shiftdown(*heap, 0)
		} else {
			*heap = append(*heap, num)
			shiftup(*heap, len(*heap)-1)
		}
	}

	for _, num := range nums {
		push(&heap, num)
	}
	return top(heap)
}
