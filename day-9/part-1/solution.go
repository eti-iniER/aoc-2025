package main

import (
	"aoc-2025/day-9/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	redTiles := common.ExtractRedTiles(input)
	n := len(redTiles)
	maxArea := 0

	for i := range n {
		for j := i + 1; j < n; j++ {
			area := common.FindArea(redTiles[i], redTiles[j])
			maxArea = max(area, maxArea)
		}
	}

	fmt.Printf("The maximum area achievable is %d\n", maxArea)
}
