package common

func ToByteGrid(input []string) [][]byte {
	n, m := len(input), len(input[0])

	grid := make([][]byte, n)
	for i := range n {
		grid[i] = make([]byte, m)

		for j := range m {
			grid[i][j] = input[i][j]
		}
	}

	return grid
}
