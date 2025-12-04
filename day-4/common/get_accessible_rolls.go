package common

func GetAccessibleRolls(grid [][]byte) [][]int {
	directions := [][]int{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}
	n, m := len(grid), len(grid[0])

	canBeAccessed := [][]int{}

	for i := range n {
		for j := range m {
			if grid[i][j] != '@' {
				continue
			}

			var surroundingRollCount int

			for _, d := range directions {
				di, dj := d[0], d[1]
				ni, nj := i+di, j+dj

				if 0 <= ni && ni < n && 0 <= nj && nj < m {
					if grid[ni][nj] == '@' {
						surroundingRollCount++
					}
				}
			}

			if surroundingRollCount < 4 {
				canBeAccessed = append(canBeAccessed, []int{i, j})
			}
		}
	}

	return canBeAccessed
}
