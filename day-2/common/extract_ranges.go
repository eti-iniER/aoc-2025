package common

import (
	"strconv"
	"strings"
)

func toInt(s string) int {
	value, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return value
}

func ExtractRanges(line string) [][]int {
	rangeStrings := strings.Split(line, ",")
	var ranges [][]int

	for _, rangeString := range rangeStrings {
		rangeValues := strings.Split(rangeString, "-")
		start := toInt(rangeValues[0])
		end := toInt(rangeValues[1])

		thisRange := []int{start, end}
		ranges = append(ranges, thisRange)
	}

	return ranges
}
