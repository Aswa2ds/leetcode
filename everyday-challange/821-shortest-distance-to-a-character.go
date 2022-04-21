package everydaychallange

func shortestToChar(s string, c byte) []int {
	length := len(s)
	leftIndex := -length
	rightIndex := 2 * length
	resultDistance := make([]int, length)

	for i := 0; i < length; i++ {
		if s[i] == c {
			leftIndex = i
		}
		resultDistance[i] = i - leftIndex
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := length - 1; i >= 0; i-- {
		if s[i] == c {
			rightIndex = i
		}
		resultDistance[i] = min(rightIndex-i, resultDistance[i])
	}

	return resultDistance
}
