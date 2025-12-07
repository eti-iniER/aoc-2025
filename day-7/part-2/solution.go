package main

import (
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	rows, columns := len(input), len(input[0])
	tachyonBeamCount := make([]int, columns)
	timelineCount := 0

	for i := range rows {
		for j := range columns {
			switch input[i][j] {
			case '^':
				if tachyonBeamCount[j] > 0 {
					timelineCount += tachyonBeamCount[j]

					if j-1 >= 0 && input[i][j-1] == '.' {
						tachyonBeamCount[j-1] += tachyonBeamCount[j]
					}
					if j+1 < columns && input[i][j+1] == '.' {
						tachyonBeamCount[j+1] += tachyonBeamCount[j]
					}
				}
				tachyonBeamCount[j] = 0
			case 'S':
				tachyonBeamCount[j]++
			default:

			}
		}
	}

	fmt.Printf("There are exactly %d timeline(s)\n", timelineCount+1)
}
