package main

import (
	"aoc-2025/day-2/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	ranges := common.ExtractRanges(input[0])
	var total int

	for _, r := range ranges {
		start, end := r[0], r[1]

		for i := start; i <= end; i++ {
			if common.IsInvalid(i, true) {
				total += i
			}
		}
	}

	fmt.Printf("The sum of invalid IDs is %d", total)
}
