package everydaychallange

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])

	pacific := make([][]int, m)
	atlantic := make([][]int, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]int, n)
		atlantic[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pacific[i][j] = -1
			atlantic[i][j] = -1
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(ocean [][]int, i, j int)
	dfs = func(ocean [][]int, i, j int) {
		if ocean[i][j] != -1 {
			return
		}
		ocean[i][j] = heights[i][j]
		for _, direction := range directions {
			newI, newJ := i+direction[0], j+direction[1]
			if newI < 0 || newI >= m || newJ < 0 || newJ >= n || ocean[newI][newJ] != -1 || heights[newI][newJ] < heights[i][j] {
				continue
			}
			dfs(ocean, newI, newJ)
		}
	}

	for j := 0; j < n; j++ {
		dfs(pacific, 0, j)
		dfs(atlantic, m-1, j)
	}

	for i := 0; i < m; i++ {
		dfs(pacific, i, 0)
		dfs(atlantic, i, n-1)
	}

	result := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] != -1 && atlantic[i][j] != -1 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
