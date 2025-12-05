package main

import (
	"aoc-2025/day-5/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	ranges, ingredients := common.ExtractIngredientData(input)
	freshCount := 0

	for _, ingredient := range ingredients {
		isSpoiled := true

		for _, r := range ranges {
			if r[0] <= ingredient && r[1] >= ingredient {
				isSpoiled = false
				break
			}
		}

		if !isSpoiled {
			freshCount++
		}
	}

	fmt.Printf("There are %d fresh ingredient(s)", freshCount)
}
