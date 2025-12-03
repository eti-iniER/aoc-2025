package main

import (
	"aoc-2025/day-3/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	var total int

	for _, line := range input {
		value := common.MaxFromLine(line, 2)
		total += value
	}

	fmt.Printf("The total value is: %d", total)
}
