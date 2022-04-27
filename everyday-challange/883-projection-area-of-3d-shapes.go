package everydaychallange

func projectionArea(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	column := 0
	for i := 0; i < len(grid); i++ {
		column = max(column, len(grid[i]))
	}

	front, left, up := 0, 0, 0
	frontStatics := make([]int, column)

	for i := 0; i < len(grid); i++ {
		highest := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				up++
				highest = max(highest, grid[i][j])
				frontStatics[j] = max(frontStatics[j], grid[i][j])
			}
		}
		left += highest
	}

	for _, v := range frontStatics {
		front += v
	}

	return front + left + up
}
