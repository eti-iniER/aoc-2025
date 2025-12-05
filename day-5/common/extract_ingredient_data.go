package common

import (
	"aoc-2025/utils"
	"strings"
)

func ExtractIngredientData(input []string) ([][]int, []int) {
	ranges := []string{}
	ingredients := []string{}
	hasSeenBlankLine := false

	for _, line := range input {
		if line == "" {
			hasSeenBlankLine = true
			continue
		}

		if hasSeenBlankLine {
			ingredients = append(ingredients, line)
		} else {
			ranges = append(ranges, line)
		}
	}

	intRanges := [][]int{}
	intIngredients := []int{}

	for _, r := range ranges {
		ingredientRange := strings.Split(r, "-")
		intRange := []int{
			utils.ToInt(ingredientRange[0]), utils.ToInt(ingredientRange[1]),
		}
		intRanges = append(intRanges, intRange)
	}

	for _, i := range ingredients {
		intIngredients = append(intIngredients, utils.ToInt(i))
	}

	return intRanges, intIngredients
}
