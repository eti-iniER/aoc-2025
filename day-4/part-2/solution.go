package main

import (
	"aoc-2025/day-4/common"
	"aoc-2025/utils"
	"fmt"
)

func removeAccessibleRolls(accessibleRolls [][]int, grid [][]byte) {
	for _, coord := range accessibleRolls {
		i, j := coord[0], coord[1]

		grid[i][j] = '.'
	}
}

func main() {
	input := utils.ReadInput("../input.txt", false)
	grid := common.ToByteGrid(input)

	accessibleRolls := common.GetAccessibleRolls(grid)
	canBeRemoved := 0

	for len(accessibleRolls) > 0 {
		canBeRemoved += len(accessibleRolls)
		removeAccessibleRolls(accessibleRolls, grid)
		accessibleRolls = common.GetAccessibleRolls(grid)
	}

	fmt.Printf("The number of rolls that can be accessed is %d", canBeRemoved)
}
