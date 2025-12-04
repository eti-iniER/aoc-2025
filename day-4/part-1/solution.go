package main

import (
	"aoc-2025/day-4/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	grid := common.ToByteGrid(input)
	accessibleRolls := common.GetAccessibleRolls(grid)

	fmt.Printf("The number of rolls that can be accessed is %d", len(accessibleRolls))
}
