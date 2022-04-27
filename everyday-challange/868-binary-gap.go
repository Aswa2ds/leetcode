package everydaychallange

func binaryGap(n int) int {
	lastIndex := -1
	maxGap := 0
	bitIndex := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for n > 0 {
		currBit := n % 2
		if currBit == 1 {
			if lastIndex != -1 {
				maxGap = max(bitIndex-lastIndex, maxGap)
			}
			lastIndex = bitIndex
		}
		n >>= 1
		bitIndex += 1
	}
	return maxGap
}
