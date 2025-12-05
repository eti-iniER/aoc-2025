package main

import (
	"aoc-2025/day-5/common"
	"aoc-2025/utils"
	"cmp"
	"fmt"
	"slices"
)

func comp(a, b []int) int {
	firstEqual := cmp.Compare(a[0], b[0])

	if firstEqual == 0 {
		return cmp.Compare(a[1], b[1])
	}

	return firstEqual
}

func mergeRanges(ranges [][]int) [][]int {
	mergedRanges := [][]int{}
	slices.SortFunc(ranges, comp)

	for _, r := range ranges {
		for len(mergedRanges) > 0 && mergedRanges[len(mergedRanges)-1][1] >= r[0] {
			last := mergedRanges[len(mergedRanges)-1]
			r = []int{
				min(r[0], last[0]), max(r[1], last[1]),
			}
			mergedRanges = mergedRanges[:len(mergedRanges)-1]
		}
		mergedRanges = append(mergedRanges, r)
	}

	return mergedRanges

}

func main() {
	input := utils.ReadInput("../input.txt", false)
	ranges, _ := common.ExtractIngredientData(input)
	ranges = mergeRanges(ranges)
	freshCount := 0

	for _, r := range ranges {
		freshCount += (r[1] - r[0] + 1)
	}

	fmt.Printf("There are %d fresh ingredient(s)", freshCount)
}
