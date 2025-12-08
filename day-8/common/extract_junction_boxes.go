package common

import (
	"aoc-2025/utils"
	"strings"
)

func ExtractJunctionBoxes(input []string) []JunctionBox {
	junctionBoxes := []JunctionBox{}

	for _, line := range input {
		values := strings.Split(line, ",")

		junctionBoxes = append(junctionBoxes, JunctionBox{
			utils.ToInt(values[0]), utils.ToInt(values[1]), utils.ToInt(values[2]),
		})

	}

	return junctionBoxes
}
