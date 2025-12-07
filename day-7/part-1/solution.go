package main

import (
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	rows, columns := len(input), len(input[0])
	hasTachyonBeam := make([]bool, columns)
	splitCount := 0

	for i := range rows {
		for j := range columns {
			switch input[i][j] {
			case '^':
				if hasTachyonBeam[j] {
					splitCount++

					if j-1 >= 0 && input[i][j-1] == '.' {
						hasTachyonBeam[j-1] = true
					}
					if j+1 < columns && input[i][j+1] == '.' {
						hasTachyonBeam[j+1] = true
					}
				}
				hasTachyonBeam[j] = false
			case 'S':
				hasTachyonBeam[j] = true
			default:

			}
		}
	}

	fmt.Printf("Tachyon beams were split exactly %d time(s)\n", splitCount)
}
