package common

import (
	"aoc-2025/utils"
	"strings"
)

func ExtractRedTiles(input []string) [][]int {
	redTiles := [][]int{}

	for _, line := range input {
		coords := strings.Split(line, ",")
		data := []int{
			utils.ToInt(coords[1]), utils.ToInt(coords[0]),
		}
		redTiles = append(redTiles, data)
	}

	return redTiles
}
